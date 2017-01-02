package main

import (
	"fmt"
	"net/http"

	"gopkg.in/olivere/elastic.v3"

	"github.com/alternative-vote/server/authentication"
	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/elections"
	"github.com/alternative-vote/server/generated"
	"github.com/dgrijalva/jwt-go"
)

func main() {

	client := initDB()

	//Auth middleware for admin stuff
	generated.AddMiddleware(adminAuthMiddleare, `/elections/{restOfRoute:.*}`)

	//link up controllers and start the server
	generated.RouterElectionController = &elections.Controller{Client: client}
	generated.RouterAuthenticationController = &authentication.Controller{}
	generated.StartServer("127.0.0.1:8000")
}

func initDB() *elastic.Client {
	// Create a client and connect to http://192.168.2.10:9201
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(consts.DB_LOC))
	if err != nil {
		panic(err)
	}

	// client.DeleteIndex(INDEX).Do()

	//check to see if our one index exists
	exists, err := client.IndexExists(consts.INDEX).Do()
	if err != nil {
		panic(err)
	}

	if !exists {
		_, err = client.CreateIndex(consts.INDEX).Do()
		if err != nil {
			panic(err)
		}
	}
	return client
}

func adminAuthMiddleare(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	tokenString := req.Header.Get("authorization")

	if tokenString == "" {
		panic(generated.HttpError(401).Message("missing token in authorization header"))
	}

	token, _ := jwt.ParseWithClaims(tokenString, &authentication.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(generated.HttpError(401).Message(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])))
		}
		return authentication.Secret, nil
	})

	_, ok := token.Claims.(*authentication.CustomClaims)

	if !ok {
		panic(generated.HttpError(401).Message("unable to unpack token"))
	}

	if !token.Valid {
		panic(generated.HttpError(401))
	}

	//TODO: put CustomClaims data on request context

	next(res, req)
}
