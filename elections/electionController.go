package elections

import (
	"encoding/json"
	"fmt"

	"gopkg.in/olivere/elastic.v3"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
	"github.com/dgrijalva/jwt-go"
)

type Controller struct {
	Client *elastic.Client
}

func checkError(err error) {
	if err == nil {
		return
	}

	if elastic.IsNotFound(err) {
		panic(HttpError(404))
	}

	panic(err)
}

//don't have an "elections service" yet, so just using the controller as a dumping ground service style functions

//get an election from the DB by ID
func (o *Controller) getElectionById(id string) domain.Election {
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(id).
		Do()
	checkError(err)

	var election domain.Election
	err = json.Unmarshal(*results.Source, &election)
	if err != nil {
		panic(err)
	}

	return election
}

//Save an election to the DB
func (o *Controller) saveElection(election domain.Election) {
	_, err := o.Client.Index().
		Index(consts.INDEX).
		Type("election").
		Id(election.Id).
		BodyJson(election).
		Do()
	checkError(err)
}

//decode voter claims information from a jwt token
func getClaims(tokenString string) VoterClaims {
	token, _ := jwt.ParseWithClaims(tokenString, &VoterClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(HttpError(401).Message(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])))
		}
		return consts.Secret, nil
	})

	if token == nil {
		panic(HttpError(404).Message("This election has been removed or never existed."))
	}

	if !token.Valid {
		panic(HttpError(401))
	}

	claims, ok := token.Claims.(*VoterClaims)

	if !ok {
		panic(HttpError(401).Message("unable to unpack token"))
	}

	return *claims
}
