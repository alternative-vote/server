package domain

import (
	"encoding/json"

	"github.com/alternative-vote/server/generated"
)

type Election struct {
	generated.Election
	Ballots []generated.Ballot `json:"ballots"`
}

// pretty lame stuff right here
type election struct {
	Ballots []generated.Ballot `json:"ballots"`
}

func (o *Election) UnmarshalJSON(data []byte) error {
	var err error

	// call the unmarshal function on the embedded struct
	err = o.Election.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	//now unmarshal just the ballots using the private object and then copy that value over to our public struct
	privateObject := new(election)
	err = json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	o.Ballots = privateObject.Ballots

	return nil
}
