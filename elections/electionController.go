package elections

import (
	. "github.com/alternative-vote/server/generated"
)

type Controller struct {
}

func (o *Controller) ListElections(req *ListElectionsRequest) *ListElectionsResponse {
	panic(HttpError(418))
}
func (o *Controller) CreateElection(req *CreateElectionRequest) *CreateElectionResponse {
	panic(HttpError(418))
}
func (o *Controller) GetElection(req *GetElectionRequest) *GetElectionResponse {
	panic(HttpError(418))
}
func (o *Controller) UpdateElection(req *UpdateElectionRequest) *UpdateElectionResponse {
	panic(HttpError(418))
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
