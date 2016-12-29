package elections

import (
	"encoding/json"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) GetElection(req *GetElectionRequest) *GetElectionResponse {
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		Do()
	checkError(err)

	var election domain.Election
	json.Unmarshal(*results.Source, &election)

	return &GetElectionResponse{
		StatusCode: 200,
		Body:       election.Election,
	}

}
