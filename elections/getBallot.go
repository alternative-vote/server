package elections

import . "github.com/alternative-vote/server/generated"
import "fmt"

func (o *Controller) GetBallot(req *GetBallotRequest) *GetBallotResponse {
	//Get voter claims off of the token
	claims := o.getClaims(req.PathParams.Token)

	//get the full election from the db
	election := o.getElectionById(claims.ElectionId)

	//if the ballot already exists, we return it, otherwise, create and then return it
	//this is all in memory right now - should definitely happen through db stuff
	ballotId := fmt.Sprintf("%v:%v", claims.ElectionId, claims.EmailAddress)
	ballot := o.getBallot(ballotId)
	if ballot == nil {
		ballot = &Ballot{
			Id:         ballotId,
			Voter:      claims.EmailAddress,
			ElectionId: claims.ElectionId,
		}
		o.saveBallot(*ballot)
	}

	election.Voters = nil //don't send down the list of voters for the ballot view
	return &GetBallotResponse{
		StatusCode: 200,
		Body: GetBallotResponseBody{
			Election: election,
			Ballot:   *ballot,
		},
	}
}
