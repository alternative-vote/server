package elections

import (
	"github.com/alternative-vote/server/consts"
	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) DeleteElection(req *DeleteElectionRequest) *DeleteElectionResponse {
	_, err := o.Client.Delete().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		Do()
	checkError(err)

	return &DeleteElectionResponse{
		StatusCode: 200,
	}
}
