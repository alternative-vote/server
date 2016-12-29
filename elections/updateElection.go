package elections

import (
	"encoding/json"
	"reflect"
	"strings"

	"fmt"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
	"github.com/davecgh/go-spew/spew"
)

func (o *Controller) UpdateElection(req *UpdateElectionRequest) *UpdateElectionResponse {
	var wireElection domain.Election
	var dbElection domain.Election

	//pull the election off the wire
	wireElection.Election = req.Body

	//get this election from the DB
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		Do()
	checkError(err)
	json.Unmarshal(*results.Source, &dbElection)

	changedValues := Extend(&dbElection, &wireElection, TitleArray(wireElection.MetaData.GetDeserializedProperties()))

	spew.Dump(changedValues)
	// election.DateUpdated.Time = time.Now().UTC()
	// _, err := o.Client.Index().
	// 	Index(consts.INDEX).
	// 	Type("election").
	// 	Id(req.PathParams.Id).
	// 	BodyJson(election).
	// 	Do()

	// checkError(err)
	panic(HttpError(418))

	// return &UpdateElectionResponse{
	// 	StatusCode: 200,
	// 	Body:       election.Election,
	// }
}

func Extend(old interface{}, new interface{}, properties []string) []ChangedValue {
	ret := []ChangedValue{}

	rOld := reflect.ValueOf(old).Elem()
	rNew := reflect.ValueOf(new).Elem()
	for _, propName := range properties {
		fmt.Println("working against property: ", propName)
		rOldProp := rOld.FieldByName(propName)
		rNewProp := rNew.FieldByName(propName)

		spew.Dump(rOldProp.Kind().String())
		spew.Dump(rNewProp.Kind().String())

		// if rOldProp.IsNil() && rNewProp.IsNil() {
		// 	continue
		// }

		// if rOldProp.IsNil() != rNewProp.IsNil() || rOldProp.Elem().Interface() != rNewProp.Elem().Interface() {
		// 	var to, from interface{}
		// 	if !rOldProp.IsNil() {
		// 		from = rOldProp.Elem().Interface()
		// 	}
		// 	if !rNewProp.IsNil() {
		// 		to = rNewProp.Elem().Interface()
		// 	}

		// 	ret = append(ret, ChangedValue{propName, from, to})

		// 	rOldProp.Set(rNewProp)
		// }

	}
	return ret
}

type ChangedValue struct {
	Name string
	From interface{}
	To   interface{}
}

func TitleArray(ss []string) []string {
	ret := []string{}
	for _, v := range ss {
		s := v

		ret = append(ret, strings.Title(s))
	}
	return ret
}
