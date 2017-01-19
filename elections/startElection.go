package elections

import (
	"encoding/json"
	"time"

	"github.com/alternative-vote/server/consts"
	
	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) StartElection(req *StartElectionRequest) *StartElectionResponse {
	var election Election

	//get this election from the DB
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		Do()
	checkError(err)
	json.Unmarshal(*results.Source, &election)

	//if we are alrleady running, then early return
	if election.State == consts.Running {
		return &StartElectionResponse{
			StatusCode: 200,
			Body:       election,
		}
	}

	//if this is a complete election, then it's a 409
	if election.State == consts.Complete {
		panic(HttpError(409).Message("can't start an election that's complete"))
	}

	//otherwise, let's change the statue of this election
	election.State = consts.Running
	election.DateUpdated.Time = time.Now().UTC()
	election.DateStarted.Time = time.Now().UTC()

	//if that worked, let's send out emails to the voters
	for i, voter := range election.Voters {
		go sendEmail(election, voter.Email)
		election.Voters[i].EmailSent = true
	}

	//save changes to the db
	o.saveElection(election)

	return &StartElectionResponse{
		StatusCode: 200,
		Body:       election,
	}
}
