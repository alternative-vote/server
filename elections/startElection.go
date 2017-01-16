package elections

import (
	"encoding/json"
	"time"

	"fmt"

	"github.com/alternative-vote/server/consts"
	"github.com/alternative-vote/server/domain"
	. "github.com/alternative-vote/server/generated"
	"github.com/go-gomail/gomail"
)

func (o *Controller) StartElection(req *StartElectionRequest) *StartElectionResponse {
	var election domain.Election

	//get this election from the DB
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		Do()
	checkError(err)
	json.Unmarshal(*results.Source, &election)

	// //if we are alrleady running, then early return
	// if election.State == consts.Running {
	// 	return &StartElectionResponse{
	// 		StatusCode: 200,
	// 		Body:       election.Election,
	// 	}
	// }

	// //if this is a complete election, then it's a 409
	// if election.State == consts.Complete {
	// 	panic(HttpError(409).Message("can't start an election that's complete"))
	// }

	//otherwise, let's change the statue of this election
	election.State = consts.Running
	election.DateUpdated.Time = time.Now().UTC()
	election.DateStarted.Time = time.Now().UTC()

	//save changes to the db
	_, err = o.Client.Index().
		Index(consts.INDEX).
		Type("election").
		Id(req.PathParams.Id).
		BodyJson(election).
		Do()
	checkError(err)

	//if that worked, let's send out emails to the voters
	for _, emailAddress := range election.Voters {
		go sendEmail(election, emailAddress)
	}

	
	return &StartElectionResponse{
		StatusCode: 200,
		Body:       election.Election,
	}
}

func sendEmail(election domain.Election, emailAddress string) {

	token := GetVoterToken(election.Id, emailAddress)

	m := gomail.NewMessage()
	m.SetHeader("From", "electioneer.io@gmail.com")
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", "Electioneer says it's time to vote!")
	m.SetBody("text/html", fmt.Sprintf(`
    This sure is an email.
    <br />
    <br />
    <a target=_blank href="https://electioneer.io/vote/%v">click here to vote</a>
    `, token))

	d := gomail.NewDialer("smtp.gmail.com", 587, "electioneer.io@gmail.com", "lawl1234")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("email sent to ", emailAddress)
	fmt.Println("token = ", token)
}
