package elections

import (
	"fmt"

	"github.com/alternative-vote/server/consts"
	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) UpdateBallot(req *UpdateBallotRequest) *UpdateBallotResponse {
	var claims VoterClaims
	//TEMP - generate a random voter claims to make testing easier in dev
	// if req.PathParams.Token == "test" {
	// 	claims = VoterClaims{
	// 		EmailAddress: uuid.NewV4().String() + "@fake.com",
	// 		ElectionId:   TestElectionID,
	// 	}
	// } else {
	//Get voter claims off of the token
	claims = o.getClaims(req.PathParams.Token)
	// }

	//get the full election from the db
	election := o.getElectionById(claims.ElectionId)

	if election.State != consts.Running {
		panic(HttpError(409).Message("election must be running in order to submit a ballot"))
	}

	ballot := req.Body
	ballot.Id = fmt.Sprintf("%v:%v", claims.ElectionId, claims.EmailAddress)
	ballot.Voter = claims.EmailAddress
	ballot.ElectionId = claims.ElectionId

	//stop this request if the ballot has already been sumbitted
	dbBallot := o.getBallot(ballot.Id)

	if dbBallot != nil && dbBallot.IsSubmitted {
		panic(HttpError(409).Message("ballot has already been submitted"))
	}

	o.saveBallot(ballot)

	election.Voters = nil //don't send down the list of voters for the ballot view
	return &UpdateBallotResponse{
		StatusCode: 200,
		Body: UpdateBallotResponseBody{
			Election: election,
			Ballot:   ballot,
		},
	}

}
