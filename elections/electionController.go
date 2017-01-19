package elections

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gopkg.in/olivere/elastic.v3"

	"strings"
	"time"

	"github.com/alternative-vote/server/consts"
	"github.com/dgrijalva/jwt-go"

	. "github.com/alternative-vote/server/generated"
	"github.com/go-gomail/gomail"
)

type Controller struct {
	Client *elastic.Client
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

//don't have an "elections service" yet, so just using the controller as a dumping ground service style functions

//get an election from the DB by ID
func (o *Controller) getElectionById(id string) Election {
	fmt.Println("getElectionById called")
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("election").
		Id(id).
		Do()

	if elastic.IsNotFound(err) {
		panic(HttpError(404))
	}

	//deal with db rate limiting
	if err != nil && strings.Contains(err.Error(), "429") {
		fmt.Println("Slowing down....")
		time.Sleep(time.Millisecond * 50)
		return o.getElectionById(id)
	}

	//something bad happened
	if err != nil {
		fmt.Println("error in getElectionById:")
		panic(err)
	}

	var election Election
	err = json.Unmarshal(*results.Source, &election)

	return election
}

//Save an election to the DB
func (o *Controller) saveElection(election Election) {
	fmt.Println("saveElection called")
	_, err := o.Client.Index().
		Index(consts.INDEX).
		Type("election").
		Id(election.Id).
		BodyJson(election).
		Do()

	//deal with db rate limiting
	if err != nil && strings.Contains(err.Error(), "429") {
		fmt.Println("Slowing down....")
		time.Sleep(time.Millisecond * 50)
		o.saveElection(election)
		return
	}

	//something bad happened
	if err != nil {
		fmt.Println("error in saveElection:")
		panic(err)
	}
}

func (o *Controller) getBallot(id string) *Ballot {
	results, err := o.Client.Get().
		Index(consts.INDEX).
		Type("ballot").
		Id(id).
		Do()

	if elastic.IsNotFound(err) {
		return nil
	}

	//deal with db rate limiting
	if err != nil && strings.Contains(err.Error(), "429") {
		fmt.Println("Slowing down....")
		time.Sleep(time.Millisecond * 50)
		return o.getBallot(id)
	}

	//something bad happened
	if err != nil {
		fmt.Println("error in getElectionById:")
		panic(err)
	}

	var ballot Ballot
	err = json.Unmarshal(*results.Source, &ballot)

	return &ballot
}

func (o *Controller) saveBallot(ballot Ballot) {
	_, err := o.Client.Index().
		Index(consts.INDEX).
		Type("ballot").
		Id(ballot.Id).
		BodyJson(ballot).
		Do()

	//deal with db rate limiting
	if err != nil && strings.Contains(err.Error(), "429") {
		fmt.Println("Slowing down....")
		time.Sleep(time.Millisecond * 50)
		o.saveBallot(ballot)
		return
	}

	//something bad happened
	if err != nil {
		fmt.Println("error in saveBallot:")
		panic(err)
	}
}

func (o *Controller) getBallots(electionId string) []Ballot {
	ballots := make([]Ballot, 0)

	query := elastic.NewSimpleQueryStringQuery(electionId)
	searchResults, err := o.Client.
		Search(consts.INDEX).
		Query(query).
		Type("ballot").
		From(0).
		Size(10000).
		Do()
	checkError(err)

	//deal with db rate limiting
	if err != nil && strings.Contains(err.Error(), "429") {
		fmt.Println("Slowing down....")
		time.Sleep(time.Millisecond * 50)
		return o.getBallots(electionId)
	}

	//something bad happened
	if err != nil {
		fmt.Println("error in saveBallot:")
		panic(err)
	}

	var b Ballot
	for _, item := range searchResults.Each(reflect.TypeOf(b)) {
		if ballot, ok := item.(Ballot); ok && ballot.ElectionId == electionId {
			ballots = append(ballots, ballot)
		}
	}

	return ballots
}

//decode voter claims information from a jwt token
func getClaims(tokenString string) VoterClaims {
	token, _ := jwt.ParseWithClaims(tokenString, &VoterClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(HttpError(401).Message(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])))
		}
		return consts.Secret, nil
	})

	if token == nil {
		panic(HttpError(404).Message("This election has been removed or never existed."))
	}

	if !token.Valid {
		panic(HttpError(401))
	}

	claims, ok := token.Claims.(*VoterClaims)

	if !ok {
		panic(HttpError(401).Message("unable to unpack token"))
	}

	return *claims
}

func sendEmail(election Election, emailAddress string) {

	token := GetVoterToken(election.Id, emailAddress)

	m := gomail.NewMessage()
	m.SetHeader("From", "electioneer.io@gmail.com")
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", fmt.Sprintf("Time to cast your vote in %v!", election.Title))
	m.SetBody("text/html", fmt.Sprintf(`

%v

<div class="body">
<p>You have been invited to vote in an election!</p>
<div class="card">
    <h1>%v</h1>
    <h2>%v</h2>
    <p>
    <a target=_blank href="https://electioneer.herokuapp.com/vote/%v"><button>Go Vote!</button></a>
    </p>
</div>
<p>
This link acts as your voter registration,  so don't share it with anyone else!
</p>
</div>
    `, css, election.Title, election.Subtitle, token))

	d := gomail.NewDialer("smtp.gmail.com", 587, "logrhythm.hackathon@gmail.com", "lawl1234")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("email sent to ", emailAddress)
	fmt.Println("token = ", token)
}

func sendResultsEmail(election Election, emailAddress string) {

	token := GetVoterToken(election.Id, emailAddress)

	m := gomail.NewMessage()
	m.SetHeader("From", "electioneer.io@gmail.com")
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", fmt.Sprintf("The results are in for %v!", election.Title))
	m.SetBody("text/html", fmt.Sprintf(`

%v

<div class="body">
<div class="card">
    <h1>%v</h1>
    <h2>%v</h2>
    <p>
    <a target=_blank href="https://electioneer.herokuapp.com/vote/%v"><button>See Results!</button></a>
    </p>
</div>

</div>
    `, css, election.Title, election.Subtitle, token))

	d := gomail.NewDialer("smtp.gmail.com", 587, "logrhythm.hackathon@gmail.com", "lawl1234")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("results email sent to ", emailAddress)
	fmt.Println("token = ", token)
}

const css = `
		<style type="text/css">
		.body {
			padding: 20px;
			font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", "Helvetica", "Arial", sans-serif;
		}
		.card {
			margin: auto 0;
			max-width: 600px;
			padding: 20px;
			box-shadow: 0 2px 3px rgba(10, 10, 10, 0.1), 0 0 0 1px rgba(10, 10, 10, 0.1);
			text-align: center;
		}

		button {
			-moz-appearance: none;
			-webkit-appearance: none;
			-webkit-box-align: center;
			-ms-flex-align: center;
			align-items: center;
			border: none;
			border-radius: 3px;
			box-shadow: none;
			display: -webkit-inline-box;
			display: -ms-inline-flexbox;
			display: inline-flex;
			font-size: 1rem;
			height: 2.285em;
			-webkit-box-pack: start;
			-ms-flex-pack: start;
			justify-content: flex-start;
			line-height: 1.5;
			padding-left: 0.75em;
			padding-right: 0.75em;
			position: relative;
			vertical-align: top;
			-webkit-touch-callout: none;
			-webkit-user-select: none;
			-moz-user-select: none;
			-ms-user-select: none;
			user-select: none;
			background-color: white;
			border: 1px solid #dbdbdb;
			color: #363636;
			cursor: pointer;
			-webkit-box-pack: center;
			-ms-flex-pack: center;
			justify-content: center;
			padding-left: 0.75em;
			padding-right: 0.75em;
			text-align: center;
			white-space: nowrap;
		}

		a {
			text-decoration: none;
		}
		</style>`
