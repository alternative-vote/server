package generated

import "fmt"

func (o *ElectionResultsStats) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"start", "end", "numVoters", "ballotsSubmitted", "averageCandidatesRanked"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *ElectionResultsStats) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "start") {
		//start is a primative
		startErr := func(propValue APITime, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid date-time
    if FormatValidators["date-time"] != nil {
      valid, formatError := FormatValidators["date-time"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.start: %v", parentName, formatError))
      }
    }





    return ret
}(o.Start, parentName)
		if startErr != nil {
			errors = append(errors, startErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "end") {
		//end is a primative
		endErr := func(propValue APITime, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid date-time
    if FormatValidators["date-time"] != nil {
      valid, formatError := FormatValidators["date-time"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.end: %v", parentName, formatError))
      }
    }





    return ret
}(o.End, parentName)
		if endErr != nil {
			errors = append(errors, endErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "numVoters") {
		//numVoters is a primative
		numVotersErr := func(propValue int64, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.NumVoters, parentName)
		if numVotersErr != nil {
			errors = append(errors, numVotersErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "ballotsSubmitted") {
		//ballotsSubmitted is a primative
		ballotsSubmittedErr := func(propValue int64, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.BallotsSubmitted, parentName)
		if ballotsSubmittedErr != nil {
			errors = append(errors, ballotsSubmittedErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "averageCandidatesRanked") {
		//averageCandidatesRanked is a primative
		averageCandidatesRankedErr := func(propValue float64, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.AverageCandidatesRanked, parentName)
		if averageCandidatesRankedErr != nil {
			errors = append(errors, averageCandidatesRankedErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



    return errors
}

func (o *ElectionResults) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"orderedCandidates", "stats", "fullData"}, fieldFromJSON) {
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
		//orderedCandidates is an array of structs
		for _, v := range o.OrderedCandidates {
			errors = append(errors, v.validate(parentName + ".orderedCandidates")...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "stats") {
		//stats is a struct
		errors = append(errors, o.Stats.validate(parentName + ".stats")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.Stats.validate(parentName + ".stats")...)



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "fullData") {
		//fullData is an array of primatives
		for _, v := range o.FullData {
			fullDataErr := func(propValue interface{}, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(v, parentName)
			if fullDataErr != nil {
				errors = append(errors, fullDataErr...)
			}
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



    return errors
}

func (o *Election) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"id", "dateCreated", "dateUpdated", "title", "subtitle", "description", "startDate", "endDate", "dateStarted", "dateEnded", "state", "owner", "voters", "candidates", "results"}, fieldFromJSON) {
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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "startDate") {
		//startDate is a struct
		errors = append(errors, o.StartDate.validate(parentName + ".startDate")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.StartDate.validate(parentName + ".startDate")...)



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "endDate") {
		//endDate is a struct
		errors = append(errors, o.EndDate.validate(parentName + ".endDate")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.EndDate.validate(parentName + ".endDate")...)



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "owner") {
		//owner is a struct
		errors = append(errors, o.Owner.validate(parentName + ".owner")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.Owner.validate(parentName + ".owner")...)



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "voters") {
		//voters is an array of structs
		for _, v := range o.Voters {
			errors = append(errors, v.validate(parentName + ".voters")...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "candidates") {
		//candidates is an array of structs
		for _, v := range o.Candidates {
			errors = append(errors, v.validate(parentName + ".candidates")...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "results") {
		//results is a struct
		errors = append(errors, o.Results.validate(parentName + ".results")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.Results.validate(parentName + ".results")...)



    return errors
}

func (o *Voter) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"email", "voteEmailSent", "resultsEmailSent"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *Voter) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "voteEmailSent") {
		//voteEmailSent is a primative
		voteEmailSentErr := func(propValue bool, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.VoteEmailSent, parentName)
		if voteEmailSentErr != nil {
			errors = append(errors, voteEmailSentErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "resultsEmailSent") {
		//resultsEmailSent is a primative
		resultsEmailSentErr := func(propValue bool, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.ResultsEmailSent, parentName)
		if resultsEmailSentErr != nil {
			errors = append(errors, resultsEmailSentErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'

		//set default if nothing came off the wire for this property
		if !hasElem(o.MetaData.GetDeserializedProperties(), "manual") {
			o.Manual = func() bool {
    return true
}()
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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



    return errors
}

func (o *Ballot) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"id", "electionId", "voter", "votes", "isSubmitted"}, fieldFromJSON) {
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
	if hasElem(o.MetaData.GetDeserializedProperties(), "id") {
		//id is a primative
		idErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    






    return ret
}(o.Id, parentName)
		if idErr != nil {
			errors = append(errors, idErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "electionId") {
		//electionId is a primative
		electionIdErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid uuid
    if FormatValidators["uuid"] != nil {
      valid, formatError := FormatValidators["uuid"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.electionId: %v", parentName, formatError))
      }
    }





    return ret
}(o.ElectionId, parentName)
		if electionIdErr != nil {
			errors = append(errors, electionIdErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "voter") {
		//voter is a primative
		voterErr := func(propValue string, parentName string) []string {
    ret := []string{}
    v := propValue
    _ = &v //if there's no validation, we need to trick the compiler into thinking v is getting used
    

    //check to see if this is a valid email
    if FormatValidators["email"] != nil {
      valid, formatError := FormatValidators["email"](propValue)
      if !valid {
        errors = append(errors, fmt.Sprintf("%v.voter: %v", parentName, formatError))
      }
    }





    return ret
}(o.Voter, parentName)
		if voterErr != nil {
			errors = append(errors, voterErr...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "votes") {
		//votes is an array of structs
		for _, v := range o.Votes {
			errors = append(errors, v.validate(parentName + ".votes")...)
		}
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



    return errors
}

func (o *LoginRequestBody) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"email", "password"}, fieldFromJSON) {
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
	if !hasElem(o.MetaData.GetDeserializedProperties(), "email") {
		 errors = append(errors, fmt.Sprintf("%v.email is a required field", parentName))
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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'


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

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'



    return errors
}

func (o *GetBallotResponseBody) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"election", "ballot"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *GetBallotResponseBody) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "election") {
		//election is a struct
		errors = append(errors, o.Election.validate(parentName + ".election")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.Election.validate(parentName + ".election")...)



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "ballot") {
		//ballot is a struct
		errors = append(errors, o.Ballot.validate(parentName + ".ballot")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.Ballot.validate(parentName + ".ballot")...)



    return errors
}

func (o *UpdateBallotResponseBody) extraFields(parentName string) []string {
	ret := []string{}

    for _, fieldFromJSON := range o.MetaData.GetDeserializedProperties() {
		if !hasElem([]string{"election", "ballot"}, fieldFromJSON) {
			ret = append(ret, parentName+"."+fieldFromJSON)
		}
	}

    return ret
}

func (o *UpdateBallotResponseBody) validate(parentName string) []string {
	errors := []string{}

    //check for extra fields first
    extraFields := o.extraFields(parentName)
	if len(extraFields) > 0 {
		errors = append(errors, fmt.Sprintf("extra fields not allowed: %v", extraFields))
	}


	//go through each property

	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "election") {
		//election is a struct
		errors = append(errors, o.Election.validate(parentName + ".election")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.Election.validate(parentName + ".election")...)



	//only run validation on stuff that came over the wire
	if hasElem(o.MetaData.GetDeserializedProperties(), "ballot") {
		//ballot is a struct
		errors = append(errors, o.Ballot.validate(parentName + ".ballot")...)
	}

	//This is pretty bad - need to set defaults on embedded structs that didn't come over the wire'
	errors = append(errors, o.Ballot.validate(parentName + ".ballot")...)



    return errors
}

