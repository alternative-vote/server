package generated

import "encoding/json"

type electionresults ElectionResults

func (o *ElectionResults) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(electionresults)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = ElectionResults(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type election Election

func (o *Election) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(election)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = Election(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type timer Timer

func (o *Timer) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(timer)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = Timer(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type ballot Ballot

func (o *Ballot) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(ballot)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = Ballot(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type user User

func (o *User) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(user)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = User(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type candidate Candidate

func (o *Candidate) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(candidate)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = Candidate(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type loginrequestbody LoginRequestBody

func (o *LoginRequestBody) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(loginrequestbody)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = LoginRequestBody(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type getballotresponsebody GetBallotResponseBody

func (o *GetBallotResponseBody) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(getballotresponsebody)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = GetBallotResponseBody(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

type updateballotresponsebody UpdateBallotResponseBody

func (o *UpdateBallotResponseBody) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(updateballotresponsebody)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = UpdateBallotResponseBody(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

