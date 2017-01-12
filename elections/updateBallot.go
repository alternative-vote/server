package elections

import . "github.com/alternative-vote/server/generated"

func (o *Controller) UpdateBallot(req *UpdateBallotRequest) *UpdateBallotResponse {

	//Get voter claims off of the token
	claims := getClaims(req.PathParams.Token)

	//get the full election from the db
	election := o.getElectionById(claims.ElectionId)

	ballot := req.Body
	var ballotSaved bool
	for i := range election.Ballots {
		if election.Ballots[i].Voter == claims.EmailAddress {
			election.Ballots[i] = ballot
			ballotSaved = true
			break
		}
	}
	if !ballotSaved {
		election.Ballots = append(election.Ballots, ballot)
	}
	o.saveElection(election)

	election.Voters = []string{} //don't send down the list of voters for the ballot view
	return &UpdateBallotResponse{
		StatusCode: 200,
		Body: UpdateBallotResponseBody{
			Election: election.Election,
			Ballot:   ballot,
		},
	}

}
