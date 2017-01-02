package generated

import "fmt"

func (o *ElectionRoles) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"voters", "administrators"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *ElectionRoles) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "voters") {
		//voters is a struct
		errors = append(errors, o.Voters.validate(parentName + ".voters")...)
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "administrators") {
		//administrators is a struct
		errors = append(errors, o.Administrators.validate(parentName + ".administrators")...)
	}

    return errors
}

func (o *ElectionResults) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"orderedCandidates"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *ElectionResults) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "orderedCandidates") {
		//orderedCandidates is an array of primatives
		for _, v := range o.OrderedCandidates {
			orderedCandidatesErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(v, parentName)
			if orderedCandidatesErr != nil {
				errors = append(errors, orderedCandidatesErr...)
			}
		}
	}

    return errors
}

func (o *Election) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"id", "dateCreated", "dateUpdated", "title", "subtitle", "description", "startDate", "endDate", "dateStarted", "dateEnded", "state", "roles", "candidates", "results"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *Election) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "id") {
		//id is a primative
		idErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid uuid
    if FormatValidators["uuid"] != nil {
      valid, formatError := FormatValidators["uuid"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.id: %v", parentName, formatError))
      }
    }





    return ret
}(o.Id, parentName)
		if idErr != nil {
			errors = append(errors, idErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "dateCreated") {
		//dateCreated is a primative
		dateCreatedErr := func(propValue APITime, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid date-time
    if FormatValidators["date-time"] != nil {
      valid, formatError := FormatValidators["date-time"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.dateCreated: %v", parentName, formatError))
      }
    }





    return ret
}(o.DateCreated, parentName)
		if dateCreatedErr != nil {
			errors = append(errors, dateCreatedErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "dateUpdated") {
		//dateUpdated is a primative
		dateUpdatedErr := func(propValue APITime, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid date-time
    if FormatValidators["date-time"] != nil {
      valid, formatError := FormatValidators["date-time"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.dateUpdated: %v", parentName, formatError))
      }
    }





    return ret
}(o.DateUpdated, parentName)
		if dateUpdatedErr != nil {
			errors = append(errors, dateUpdatedErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "title") {
		//title is a primative
		titleErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Title, parentName)
		if titleErr != nil {
			errors = append(errors, titleErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "subtitle") {
		//subtitle is a primative
		subtitleErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Subtitle, parentName)
		if subtitleErr != nil {
			errors = append(errors, subtitleErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "description") {
		//description is a primative
		descriptionErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Description, parentName)
		if descriptionErr != nil {
			errors = append(errors, descriptionErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "startDate") {
		//startDate is a struct
		errors = append(errors, o.StartDate.validate(parentName + ".startDate")...)
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "endDate") {
		//endDate is a struct
		errors = append(errors, o.EndDate.validate(parentName + ".endDate")...)
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "dateStarted") {
		//dateStarted is a primative
		dateStartedErr := func(propValue APITime, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid date-time
    if FormatValidators["date-time"] != nil {
      valid, formatError := FormatValidators["date-time"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.dateStarted: %v", parentName, formatError))
      }
    }





    return ret
}(o.DateStarted, parentName)
		if dateStartedErr != nil {
			errors = append(errors, dateStartedErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "dateEnded") {
		//dateEnded is a primative
		dateEndedErr := func(propValue APITime, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid date-time
    if FormatValidators["date-time"] != nil {
      valid, formatError := FormatValidators["date-time"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.dateEnded: %v", parentName, formatError))
      }
    }





    return ret
}(o.DateEnded, parentName)
		if dateEndedErr != nil {
			errors = append(errors, dateEndedErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "state") {
		//state is a primative
		stateErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.State, parentName)
		if stateErr != nil {
			errors = append(errors, stateErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "roles") {
		//roles is a struct
		errors = append(errors, o.Roles.validate(parentName + ".roles")...)
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "candidates") {
		//candidates is an array of structs
		for _, v := range o.Candidates {
			errors = append(errors, v.validate(parentName + ".candidates")...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "results") {
		//results is a struct
		errors = append(errors, o.Results.validate(parentName + ".results")...)
	}

    return errors
}

func (o *Role) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"isPublic", "members"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *Role) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "isPublic") {
		//isPublic is a primative
		isPublicErr := func(propValue bool, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.IsPublic, parentName)
		if isPublicErr != nil {
			errors = append(errors, isPublicErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "members") {
		//members is an array of structs
		for _, v := range o.Members {
			errors = append(errors, v.validate(parentName + ".members")...)
		}
	}

    return errors
}

func (o *Timer) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"manual", "date"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *Timer) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "manual") {
		//manual is a primative
		manualErr := func(propValue bool, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Manual, parentName)
		if manualErr != nil {
			errors = append(errors, manualErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "date") {
		//date is a primative
		dateErr := func(propValue APITime, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid date-time
    if FormatValidators["date-time"] != nil {
      valid, formatError := FormatValidators["date-time"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.date: %v", parentName, formatError))
      }
    }





    return ret
}(o.Date, parentName)
		if dateErr != nil {
			errors = append(errors, dateErr...)
		}
	}

    return errors
}

func (o *Ballot) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"voter", "votes", "isSubmitted"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *Ballot) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "voter") {
		//voter is a struct
		errors = append(errors, o.Voter.validate(parentName + ".voter")...)
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "votes") {
		//votes is an array of structs
		for _, v := range o.Votes {
			errors = append(errors, v.validate(parentName + ".votes")...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "isSubmitted") {
		//isSubmitted is a primative
		isSubmittedErr := func(propValue bool, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.IsSubmitted, parentName)
		if isSubmittedErr != nil {
			errors = append(errors, isSubmittedErr...)
		}
	}

    return errors
}

func (o *User) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"id", "email", "isAccount"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *User) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "id") {
		//id is a primative
		idErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid uuid
    if FormatValidators["uuid"] != nil {
      valid, formatError := FormatValidators["uuid"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.id: %v", parentName, formatError))
      }
    }





    return ret
}(o.Id, parentName)
		if idErr != nil {
			errors = append(errors, idErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "email") {
		//email is a primative
		emailErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid email
    if FormatValidators["email"] != nil {
      valid, formatError := FormatValidators["email"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.email: %v", parentName, formatError))
      }
    }





    return ret
}(o.Email, parentName)
		if emailErr != nil {
			errors = append(errors, emailErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "isAccount") {
		//isAccount is a primative
		isAccountErr := func(propValue bool, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.IsAccount, parentName)
		if isAccountErr != nil {
			errors = append(errors, isAccountErr...)
		}
	}

    return errors
}

func (o *Candidate) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"title", "subtitle", "description"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *Candidate) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "title") {
		//title is a primative
		titleErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Title, parentName)
		if titleErr != nil {
			errors = append(errors, titleErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "subtitle") {
		//subtitle is a primative
		subtitleErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Subtitle, parentName)
		if subtitleErr != nil {
			errors = append(errors, subtitleErr...)
		}
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "description") {
		//description is a primative
		descriptionErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Description, parentName)
		if descriptionErr != nil {
			errors = append(errors, descriptionErr...)
		}
	}

    return errors
}

func (o *LoginRequestBody) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"username", "password"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *LoginRequestBody) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property
	//set check required based off of what was deserialized
	if !hasElem(o.MetaData.GetDeserializedProperties(), "username") {
		 errors = append(errors, fmt.Sprintf("%v.username is a required field", parentName))
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "username") {
		//username is a primative
		usernameErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Username, parentName)
		if usernameErr != nil {
			errors = append(errors, usernameErr...)
		}
	}
	//set check required based off of what was deserialized
	if !hasElem(o.MetaData.GetDeserializedProperties(), "password") {
		 errors = append(errors, fmt.Sprintf("%v.password is a required field", parentName))
	}

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "password") {
		//password is a primative
		passwordErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Password, parentName)
		if passwordErr != nil {
			errors = append(errors, passwordErr...)
		}
	}

    return errors
}

