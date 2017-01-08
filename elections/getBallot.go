package elections

import (
	"fmt"

	"github.com/alternative-vote/server/consts"
	. "github.com/alternative-vote/server/generated"
	"github.com/dgrijalva/jwt-go"
)

func (o *Controller) GetBallot(req *GetBallotRequest) *GetBallotResponse {

	//Get voter claims off of the token
	claims := getClaims(req.PathParams.Token)

	election := o.getElectionById(claims.ElectionId)
	election.Voters = []string{} //don't send down the list of voters for the ballot view

	/*
	   TODO: check to see if ballot exists in the DB
	   - if it doesn't, then create it
	   - if it does, then return that one
	*/
	ballot := Ballot{
		Voter: claims.EmailAddress,
	}

	return &GetBallotResponse{
		StatusCode: 200,
		Body: GetBallotResponseBody{
			Election: election.Election,
			Ballot:   ballot,
		},
	}
}

func getClaims(tokenString string) VoterClaims {
	token, _ := jwt.ParseWithClaims(tokenString, &VoterClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(HttpError(401).Message(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])))
		}
		return consts.Secret, nil
	})

	if !token.Valid {
		panic(HttpError(401))
	}

	claims, ok := token.Claims.(*VoterClaims)

	if !ok {
		panic(HttpError(401).Message("unable to unpack token"))
	}

	return *claims
}
