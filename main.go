package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"gopkg.in/olivere/elastic.v3"

	"strings"

	"path"

	"github.com/alternative-vote/server/authentication"
	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/elections"
	"github.com/alternative-vote/server/generated"
	"github.com/alternative-vote/server/utils"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	esClient := initDB()

	//link up controllers and start the server
	generated.RouterElectionController = &elections.Controller{Client: esClient}
	generated.RouterAuthenticationController = &authentication.Controller{}

	//url rewrite to strip /api - this could go away if you put /api in the swagger file
	generated.AddMiddleware(func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		req.URL.Path = strings.TrimLeft(req.URL.Path, "/api")
		if !strings.HasPrefix(req.URL.Path, "/") {
			req.URL.Path = "/" + req.URL.Path
		}
		next(res, req)
	})

	//Auth middleware for admin stuff
	generated.AddMiddleware(adminAuthMiddleare, `/elections/{restOfRoute:.*}`)

	//subrouter for all API stuff
	http.Handle("/api/", generated.Stack())

	//this will try to serve files from the client directory by matching path to filename
	//if no file is found, it will just return index.html
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		data, err := ioutil.ReadFile(path.Join("./client/", req.URL.Path))
		if err != nil && (strings.Contains(err.Error(), "no such file or directory") || strings.Contains(err.Error(), "is a directory")) {
			data, _ = ioutil.ReadFile(path.Join("./client/", "index.html"))
			http.ServeContent(res, req, "index.html", time.Now(), bytes.NewReader(data))
			return
		}

		if err != nil {
			fmt.Println(err)
			res.WriteHeader(500)
			res.Write([]byte("unknown error serving static content"))
			return
		}

		http.ServeContent(res, req, req.URL.Path, time.Now(), bytes.NewReader(data))
	})

	fmt.Println("listening at localhost:" + port)
	http.ListenAndServe(":"+port, nil)

}

func initDB() *elastic.Client {
	// Create a esClient and connect to http://192.168.2.10:9201
	esClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(consts.DB_LOC))
	if err != nil {
		panic(err)
	}

	// esClient.DeleteIndex(consts.INDEX).Do()

	//check to see if our one index exists
	exists, err := esClient.IndexExists(consts.INDEX).Do()
	if err != nil {
		panic(err)
	}

	if !exists {
		_, err = esClient.CreateIndex(consts.INDEX).Do()
		if err != nil {
			panic(err)
		}
	}
	return esClient
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
