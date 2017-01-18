package elections

import (
	. "github.com/alternative-vote/server/generated"
	"github.com/satori/go.uuid"
)

const TestElectionID = "414de022-7de8-4c59-86de-532fd4316989"

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
		claims = getClaims(req.PathParams.Token)
	}

	//get the full election from the db
	election := o.getElectionById(claims.ElectionId)

	ballot := req.Body
	ballot.Voter = claims.EmailAddress
	var ballotSaved bool
	for i := range election.Ballots {
		if election.Ballots[i].Voter == claims.EmailAddress {
			if election.Ballots[i].IsSubmitted {
				panic(HttpError(409).Message("ballot has already been submitted"))
			}

			election.Ballots[i] = ballot
			ballotSaved = true
			break
		}
	}
	if !ballotSaved {
		election.Ballots = append(election.Ballots, ballot)
		statusCode = 201
	}
	o.saveElection(election)

	election.Voters = nil //don't send down the list of voters for the ballot view
	return &UpdateBallotResponse{
		StatusCode: statusCode,
		Body: UpdateBallotResponseBody{
			Election: election.Election,
			Ballot:   ballot,
		},
	}

}
