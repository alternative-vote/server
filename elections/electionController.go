package elections

import (
	"encoding/json"

	"gopkg.in/olivere/elastic.v3"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
)

type Controller struct {
	Client *elastic.Client
}

func (o *Controller) UpdateBallot(req *UpdateBallotRequest) *UpdateBallotResponse {
	panic(HttpError(418))
}

func checkError(err error) {
	if err == nil {
		return
	}

	if elastic.IsNotFound(err) {
		panic(HttpError(404))
	}

	panic(err)
}

func (o *Controller) getElectionById(id string) domain.Election {
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(id).
		Do()
	checkError(err)

	var election domain.Election
	json.Unmarshal(*results.Source, &election)
	return election
}
