package main

import (
	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/elections"
	"github.com/alternative-vote/server/generated"
	"gopkg.in/olivere/elastic.v3"
)

func main() {

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

	// searchResults, err := client.
	// 	Search(INDEX).
	// 	Type("election").
	// 	From(0).
	// 	Size(10000).
	// 	Do()
	// if err != nil {
	// 	panic(err)
	// }

	// var e generated.Election
	// for _, item := range searchResults.Each(reflect.TypeOf(e)) {
	// 	if election, ok := item.(generated.Election); ok {
	// 		fmt.Println("Title:", election.Title)
	// 		fmt.Println("DateCreated:", election.DateCreated)
	// 		fmt.Println("DateUpdated:", election.DateUpdated)
	// 	}
	// }

	// os.Exit(0)

	// //election we are going to play with
	// testElection := generated.Election{
	// 	Id:    uuid.NewV4().String(),
	// 	Title: "test election",
	// 	Candidates: []generated.Candidate{
	// 		{Title: "A"},
	// 		{Title: "B"},
	// 	},
	// }
	// testElection.DateCreated.Time = time.Now().UTC()
	// testElection.DateUpdated.Time = time.Now().UTC()

	// //create the election
	// put, err := client.Index().
	// 	Index(INDEX).
	// 	Type("election").
	// 	Id(testElection.Id).
	// 	BodyJson(testElection).
	// 	Do()

	// if err != nil {
	// 	panic(err)
	// }
	// spew.Dump(put)

	// //get it back out
	// get, err := client.Get().
	// 	Index(INDEX).
	// 	Type("election").
	// 	Id(testElection.Id).
	// 	Do()
	// if err != nil {
	// 	panic(err)
	// }

	// spew.Dump(get)

	// var electionFromDB generated.Election
	// json.Unmarshal(*get.Source, &electionFromDB)

	// spew.Dump(electionFromDB.DateCreated)

	generated.RouterElectionController = &elections.Controller{Client: client}
	generated.StartServer("127.0.0.1:8000")
}
