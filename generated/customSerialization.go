package generated

import "encoding/json"

type electionroles ElectionRoles

func (o *ElectionRoles) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(electionroles)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = ElectionRoles(*privateObject)

	//ok, if that worked lets fill in some metadata
	for _, propertyName := range getFields(data) {
		o.MetaData.AddDeserializedProperty(propertyName)
	}

	return nil
}

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

type role Role

func (o *Role) UnmarshalJSON(data []byte) error {
	//deserialize normally (using our private dummy struct to prevent looping)
	privateObject := new(role)
	err := json.Unmarshal(data, privateObject)
	if err != nil {
		return err
	}
	*o = Role(*privateObject)

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

