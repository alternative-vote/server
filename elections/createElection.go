package elections

import (
	"time"

	"github.com/alternative-vote/server/consts"
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

	claims := utils.GetClaims(req.Context)

	var election Election

	election = req.Body

	election.Id = uuid.NewV4().String()
	election.DateCreated.Time = time.Now().UTC()
	election.DateUpdated.Time = time.Now().UTC()
	election.State = consts.Edit

	//owner comes from auth claims
	election.Owner.Email = claims.Email
	election.Owner.IsAccount = true
	election.Results.Stats.AverageCandidatesRanked = 0

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
		Body:       election,
	}
}
