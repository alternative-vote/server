package elections

import (
	"encoding/json"

	"strings"

	"github.com/alternative-vote/server/consts"

	. "github.com/alternative-vote/server/generated"
)

func (o *Controller) SendEmail(req *SendEmailRequest) *SendEmailResponse {
	var election Election

	//get this election from the DB
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		Do()
	checkError(err)
	json.Unmarshal(*results.Source, &election)

	//if we are alrleady running, then early return
	if election.State != consts.Running {
		panic(HttpError(409).
			Message("This endpoint can only be used if the election is already running.").
			Details(`
        Note that by starting an election, emails have already been sent to any
        votes on the election at that time.  This endpoint should be used to send 
        emails to anyone that was addd to the election after it started running.
        Voters that already received an email should not receive additional ones. 
        `))
	}

	//fire of a go routine to send emails one by one
	go func() {
		//if that worked, let's send out emails to the voters
		for i, voter := range election.Voters {
			if strings.ToLower(req.QueryParams.Force) == "true" || !voter.VoteEmailSent {
				o.sendEmail(election, voter.Email)
				election.Voters[i].VoteEmailSent = true
			}
		}
		//save changes to the db
		o.saveElection(election)
	}()

	return &SendEmailResponse{
		StatusCode: 200,
		Body:       election,
	}

}
