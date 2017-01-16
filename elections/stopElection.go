package elections

import (
	"time"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
	"github.com/davecgh/go-spew/spew"
)

func (o *Controller) StopElection(req *StopElectionRequest) *StopElectionResponse {
	//get this election from the DB
	election := o.getElectionById(req.PathParams.Id)
	calculateResults(&election)

	panic(HttpError(418))

	//early return if we are already complete
	if election.State == consts.Complete {
		return &StopElectionResponse{
			StatusCode: 200,
			Body:       election.Election,
		}
	}

	//if this is still in edit mode, then throw error
	if election.State == consts.Complete {
		panic(HttpError(409).Message("can't stop an election that hasn't started yet"))
	}

	//update state, metadata, and calculate results
	election.State = consts.Complete
	election.DateUpdated.Time = time.Now().UTC()
	election.DateEnded.Time = time.Now().UTC()

	calculateResults(&election)

	//save changes to the db
	o.saveElection(election)

	return &StopElectionResponse{
		StatusCode: 200,
		Body:       election.Election,
	}
}

func calculateResults(election *domain.Election) {
	candidates := []string{}
	ballots := [][]string{}

	for _, c := range election.Candidates {
		candidates = append(candidates, c.Title)
	}
	spew.Dump(election)
	for _, ballot := range election.Ballots {
		votes := []string{}
		for _, vote := range ballot.Votes {
			votes = append(votes, vote.Title)
		}
		ballots = append(ballots, votes)
	}

	spew.Dump(candidates)
	spew.Dump(ballots)

}
