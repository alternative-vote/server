package elections

import (
	"reflect"

	"github.com/alternative-vote/server/consts"
	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) ListElections(req *ListElectionsRequest) *ListElectionsResponse {
	//need to initialize to 0, otherwize an empty array will be serialized to null when writing the response
	elections := make([]Election, 0)

	searchResults, err := o.Client.
		Search(consts.INDEX).
		Type("election").
		From(0).
		Size(10000).
		Do()
	checkError(err)

	var e Election
	for _, item := range searchResults.Each(reflect.TypeOf(e)) {
		if election, ok := item.(Election); ok {

			elections = append(elections, election)
		}
	}

	return &ListElectionsResponse{
		StatusCode: 200,
		Body:       elections,
	}
}
