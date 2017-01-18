package elections

import . "github.com/alternative-vote/server/generated"

func (o *Controller) GetBallot(req *GetBallotRequest) *GetBallotResponse {
	//Get voter claims off of the token
	claims := getClaims(req.PathParams.Token)

	//get the full election from the db
	election := o.getElectionById(claims.ElectionId)

	//if the ballot already exists, we return it, otherwise, create and then return it
	//this is all in memory right now - should definitely happen through db stuff
	var ballot *Ballot
	for i := range election.Ballots {
		if election.Ballots[i].Voter == claims.EmailAddress {
			ballot = &election.Ballots[i]
			break
		}
	}
	if ballot == nil {
		ballot = &Ballot{
			Voter: claims.EmailAddress,
		}
		election.Ballots = append(election.Ballots, *ballot)
		o.saveElection(election)
	}

	election.Voters = nil //don't send down the list of voters for the ballot view
	return &GetBallotResponse{
		StatusCode: 200,
		Body: GetBallotResponseBody{
			Election: election.Election,
			Ballot:   *ballot,
		},
	}
}
