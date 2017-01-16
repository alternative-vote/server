package elections

import . "github.com/alternative-vote/server/generated"

func (o *Controller) GetElection(req *GetElectionRequest) *GetElectionResponse {

	election := o.getElectionById(req.PathParams.Id)

	
	return &GetElectionResponse{
		StatusCode: 200,
		Body:       election.Election,
	}

}
