package elections

import (
	"time"

	"github.com/alternative-vote/server/altVote"
	"github.com/alternative-vote/server/consts"

	"math"

	"fmt"

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
	// 		Body:       election,
	// 	}
	// }

	//if this is still in edit mode, then throw error
	if election.State == consts.Edit {
		panic(HttpError(409).Message("can't stop an election that hasn't started yet"))
	}

	ballots := o.getBallots(election.Id)

	//update state, metadata, and calculate results
	election.State = consts.Complete
	election.DateUpdated.Time = time.Now().UTC()
	election.DateEnded.Time = time.Now().UTC()
	election.Results = calculateResults(election, ballots)

	//Can't save NaN to elasticsearch
	if math.IsNaN(election.Results.Stats.AverageCandidatesRanked) {
		election.Results.Stats.AverageCandidatesRanked = 0
	}

	//save changes to the db
	o.saveElection(election)

	//fire of a go routine to send emails one by one
	go func() {
		for i, voter := range election.Voters {
			if !election.Voters[i].ResultsEmailSent {
				err := o.sendResultsEmail(election, voter.Email)
				if err == nil {
					election.Voters[i].ResultsEmailSent = true
				}

			}
		}
		o.saveElection(election)
	}()

	//save changes to the db again post emails sent
	o.saveElection(election)

	return &StopElectionResponse{
		StatusCode: 200,
		Body:       election,
	}
}

func calculateResults(election Election, electionBallots []Ballot) ElectionResults {
	ret := ElectionResults{}

	ret.Stats.NumVoters = int64(len(election.Voters))
	ret.Stats.Start = election.DateStarted
	ret.Stats.End = election.DateEnded

	var totalVotes float64

	candidates := []string{}
	ballots := [][]string{}

	for _, c := range election.Candidates {
		candidates = append(candidates, c.Title)
	}

	for _, ballot := range electionBallots {
		if !ballot.IsSubmitted {
			continue
		}
		votes := []string{}
		for _, vote := range ballot.Votes {
			votes = append(votes, vote.Title)
			totalVotes++
		}
		ballots = append(ballots, votes)
	}

	ret.Stats.BallotsSubmitted = int64(len(ballots))
	ret.Stats.AverageCandidatesRanked = totalVotes / float64(ret.Stats.BallotsSubmitted)

	for len(candidates) > 0 {
		//running a new election
		fmt.Printf("Running election with %v candidates and %v ballots\n", len(candidates), len(ballots))
		// spew.Dump(candidates)
		// spew.Dump(ballots[:5])
		results, err := altVote.GetResults(candidates, ballots)
		if err == altVote.NoVotes {
			fmt.Println("no more votes, we're done here.")
			break //if we get here, that means that there were some number of candidates that did not get a single vote
		}
		fmt.Printf("%v won after %v round(s)!  Removing them and rerunning...\n\n\n", results.Winner, len(results.Rounds))

		//add the winner of this election to the ordered candidates list (this is the ranked list of winners)
		ret.OrderedCandidates = append(ret.OrderedCandidates, getCandidate(election.Candidates, results.Winner))

		//now that this candidate is in the ordered list, we can remove them from the candidate list and run through the ballots again
		for i := range candidates {
			if candidates[i] == results.Winner {
				candidates = append(candidates[:i], candidates[i+1:]...)
				break
			}
		}

		//full data for this election
		ret.FullData = append(ret.FullData, results)

	}

	// spew.Dump(ret.FullData)
	ret.FullData = nil

	fmt.Println("done calculating results, orderedCandidates length = ", len(ret.OrderedCandidates))

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
