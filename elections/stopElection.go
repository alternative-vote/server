package elections

import (
	"time"

	"fmt"

	"github.com/Khelldar/altVote"
	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) StopElection(req *StopElectionRequest) *StopElectionResponse {
	//get this election from the DB
	election := o.getElectionById(req.PathParams.Id)

	//early return if we are already complete
	//TEMP - don't early return so that we can calculate results more than once for dev
	// if election.State == consts.Complete {
	// 	return &StopElectionResponse{
	// 		StatusCode: 200,
	// 		Body:       election.Election,
	// 	}
	// }

	//if this is still in edit mode, then throw error
	if election.State == consts.Edit {
		panic(HttpError(409).Message("can't stop an election that hasn't started yet"))
	}

	//update state, metadata, and calculate results
	election.State = consts.Complete
	election.DateUpdated.Time = time.Now().UTC()
	election.DateEnded.Time = time.Now().UTC()
	election.Results.OrderedCandidates = calculateResults(election)

	//save changes to the db
	o.saveElection(election)

	return &StopElectionResponse{
		StatusCode: 200,
		Body:       election.Election,
	}
}

func calculateResults(election domain.Election) []Candidate {
	ret := []Candidate{}
	candidates := []string{}
	ballots := [][]string{}

	for _, c := range election.Candidates {
		candidates = append(candidates, c.Title)
	}

	for _, ballot := range election.Ballots {
		if !ballot.IsSubmitted {
			continue
		}
		votes := []string{}
		for _, vote := range ballot.Votes {
			votes = append(votes, vote.Title)
		}
		ballots = append(ballots, votes)
	}

	for len(candidates) > 0 {
		fmt.Printf("Running election with %v candidates: %v\n", len(candidates), candidates)
		results, err := altVote.GetResults(candidates, ballots)
		if err == altVote.NoVotes {
			break //if we get here, that means that there were some number of candidates that did not get a single vote
		}
		fmt.Printf("%v won!  Removing them and rerunning...\n", results.Winner)
		ret = append(ret, getCandidate(election.Candidates, results.Winner))

		//now that this candidate is in the ordered list, we can remove them from the candidate list and run through the ballots again
		for i := range candidates {
			if candidates[i] == results.Winner {
				candidates = append(candidates[:i], candidates[i+1:]...)
				break
			}
		}
	}

	return ret
}

func getCandidate(cc []Candidate, title string) Candidate {
	for _, c := range cc {
		if c.Title == title {
			return c
		}
	}
	panic("couldn't find candidate")
}
