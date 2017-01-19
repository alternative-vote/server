package elections

import (
	"time"

	. "github.com/alternative-vote/server/generated"
	"github.com/alternative-vote/server/utils"
)

func (o *Controller) UpdateElection(req *UpdateElectionRequest) *UpdateElectionResponse {
	var wireElection Election
	var dbElection Election

	//pull the election off the wire
	wireElection = req.Body

	//get this election from the DB
	dbElection = o.getElectionById(req.PathParams.Id)

	//There's only certain fields that can be updated by the client
	updateableProps := []string{"title", "subtitle", "description", "startDate", "endDate", "candidates", "voters"}
	propsToUpdate := utils.Intersection(updateableProps, req.Body.MetaData.GetDeserializedProperties())
	utils.Extend(&dbElection, &wireElection, utils.TitleArray(propsToUpdate))

	//update timestamp
	dbElection.DateUpdated.Time = time.Now().UTC()

	//save changes to the db
	o.saveElection(dbElection)

	return &UpdateElectionResponse{
		StatusCode: 200,
		Body:       dbElection,
	}
}
