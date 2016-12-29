package elections

import (
	"time"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
	"github.com/alternative-vote/server/utils"
	"github.com/satori/go.uuid"
)

func (o *Controller) CreateElection(req *CreateElectionRequest) *CreateElectionResponse {
	if utils.Contains(req.Body.MetaData.GetDeserializedProperties(), "id") && req.Body.Id != "" {
		panic(HttpError(400).
			Message("no need to specify an 'id', it will be generated for you").
			Details("use PUT /elections/{id} if you are trying to update an existing election"))
	}

	var election domain.Election

	election.Election = req.Body

	election.Id = uuid.NewV4().String()
	election.DateCreated.Time = time.Now().UTC()
	election.DateUpdated.Time = time.Now().UTC()
	election.State = consts.Edit

	//create the election
	_, err := o.Client.Index().
		Index(consts.INDEX).
		Type("election").
		Id(election.Id).
		BodyJson(election).
		Do()

	if err != nil {
		panic(err)
	}

	return &CreateElectionResponse{
		StatusCode: 201,
		Body:       election.Election,
	}
}
