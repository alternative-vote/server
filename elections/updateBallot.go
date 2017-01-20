package elections

import (
	"fmt"

	. "github.com/alternative-vote/server/generated"
	"github.com/satori/go.uuid"
)

const TestElectionID = "03a00b8a-722b-4c23-bb62-47c46f9edd34"

func (o *Controller) UpdateBallot(req *UpdateBallotRequest) *UpdateBallotResponse {
	statusCode := 200
	var claims VoterClaims
	//TEMP - generate a random voter claims to make testing easier in dev
	if req.PathParams.Token == "test" {
		claims = VoterClaims{
			EmailAddress: uuid.NewV4().String() + "@fake.com",
			ElectionId:   TestElectionID,
		}
	} else {
		//Get voter claims off of the token
		claims = o.getClaims(req.PathParams.Token)
	}

	//get the full election from the db
	election := o.getElectionById(claims.ElectionId)

	//TEMP - I want to be able to add lots of ballots
	// if election.State != consts.Running {
	// 	panic(HttpError(409).Message("election must be running in order to submit a ballot"))
	// }

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

	// var ballotSaved bool
	// for i := range election.Ballots {
	// 	if election.Ballots[i].Voter == claims.EmailAddress {
	// 		if election.Ballots[i].IsSubmitted {
	// 			panic(HttpError(409).Message("ballot has already been submitted"))
	// 		}

	// 		election.Ballots[i] = ballot
	// 		ballotSaved = true
	// 		break
	// 	}
	// }
	// if !ballotSaved {
	// 	election.Ballots = append(election.Ballots, ballot)
	// 	statusCode = 201
	// }
	// o.saveElection(election)

	election.Voters = nil //don't send down the list of voters for the ballot view
	return &UpdateBallotResponse{
		StatusCode: statusCode,
		Body: UpdateBallotResponseBody{
			Election: election,
			Ballot:   ballot,
		},
	}

}
