package main

import (
	"encoding/json"

	"gopkg.in/olivere/elastic.v3"

	"github.com/alternative-vote/server/elections"
	"github.com/alternative-vote/server/generated"
	"github.com/davecgh/go-spew/spew"
	"github.com/satori/go.uuid"
)

const DB_LOC = "https://ifbd6dxv:6gb8n4hl99gv43ks@rowan-5667114.us-east-1.bonsaisearch.net"

func main() {

	// Create a client and connect to http://192.168.2.10:9201
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(DB_LOC))
	if err != nil {
		panic(err)
	}

	//check to see if
	exists, err := client.IndexExists("test").Do()
	if err != nil {
		panic(err)
	}

	if !exists {
		_, err = client.CreateIndex("test").Do()
		if err != nil {
			panic(err)
		}
	}

	//election we are going to play with
	testElection := generated.Election{
		Id:    uuid.NewV4().String(),
		Title: "test election",
		Candidates: []generated.Candidate{
			{Title: "A"},
			{Title: "B"},
		},
	}

	//create the election
	put, err := client.Index().
		Index("test").
		Type("election").
		Id(testElection.Id).
		BodyJson(testElection).
		Do()

	if err != nil {
		panic(err)
	}
	spew.Dump(put)

	//get it back out
	get, err := client.Get().
		Index("test").
		Type("election").
		Id(testElection.Id).
		Do()
	if err != nil {
		panic(err)
	}

	spew.Dump(get)

	var electionFromDB generated.Election
	json.Unmarshal(*get.Source, &electionFromDB)

	spew.Dump(electionFromDB)

	generated.RouterElectionController = &elections.Controller{}
	generated.StartServer("127.0.0.1:8000")
}
