package elections

import (
	"encoding/json"

	"time"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
	"github.com/alternative-vote/server/utils"
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

	//There's only certain fields that can be updated by the client
	//TODO: update roles
	updateableProps := []string{"title", "subtitle", "description", "startDate", "endDate", "candidates"}
	propsToUpdate := utils.Intersection(updateableProps, req.Body.MetaData.GetDeserializedProperties())
	utils.Extend(&dbElection, &wireElection, utils.TitleArray(propsToUpdate))

	//update timestamp
	dbElection.DateUpdated.Time = time.Now().UTC()

	//save changes to the db
	_, err = o.Client.Index().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		BodyJson(dbElection).
		Do()
	checkError(err)

	return &UpdateElectionResponse{
		StatusCode: 200,
		Body:       dbElection.Election,
	}
}
