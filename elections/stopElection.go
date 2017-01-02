package elections

import (
	"encoding/json"
	"time"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) StopElection(req *StopElectionRequest) *StopElectionResponse {
	var election domain.Election

	//get this election from the DB
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		Do()
	checkError(err)
	json.Unmarshal(*results.Source, &election)

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

	//otherwise, let's change the statue of this election
	election.State = consts.Complete
	election.DateUpdated.Time = time.Now().UTC()
	election.DateEnded.Time = time.Now().UTC()

	//save changes to the db
	_, err = o.Client.Index().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		BodyJson(election).
		Do()
	checkError(err)

	return &StopElectionResponse{
		StatusCode: 200,
		Body:       election.Election,
	}
}
