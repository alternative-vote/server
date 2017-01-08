package main

import (
	"fmt"
	"net/http"

	"gopkg.in/olivere/elastic.v3"

	"strings"

	"github.com/alternative-vote/server/authentication"
	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/elections"
	"github.com/alternative-vote/server/generated"
	"github.com/alternative-vote/server/utils"
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

	// client.DeleteIndex(consts.INDEX).Do()

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

	if !strings.HasPrefix(tokenString, "Bearer ") {
		panic(generated.HttpError(401).Message("authorization header must be a bearer token"))
	}

	tokenString = strings.Split(tokenString, " ")[1]

	token, _ := jwt.ParseWithClaims(tokenString, &authentication.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(generated.HttpError(401).Message(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])))
		}
		return consts.Secret, nil
	})

	if !token.Valid {
		panic(generated.HttpError(401))
	}

	claims, ok := token.Claims.(*authentication.CustomClaims)

	if !ok {
		panic(generated.HttpError(401).Message("unable to unpack token"))
	}

	req = req.WithContext(utils.SetClaims(req.Context(), *claims))

	next(res, req)
}
