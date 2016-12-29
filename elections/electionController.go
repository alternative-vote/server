package elections

import (
	. "github.com/alternative-vote/server/generated"
	"gopkg.in/olivere/elastic.v3"
)

type Controller struct {
	Client *elastic.Client
}

func (o *Controller) StartElection(req *StartElectionRequest) *StartElectionResponse {
	panic(HttpError(418))
}
func (o *Controller) StopElection(req *StopElectionRequest) *StopElectionResponse {
	panic(HttpError(418))
}
func (o *Controller) GetBallot(req *GetBallotRequest) *GetBallotResponse {
	panic(HttpError(418))
}
func (o *Controller) UpsertBallot(req *UpsertBallotRequest) *UpsertBallotResponse {
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
