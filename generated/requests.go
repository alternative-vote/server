package generated

import (
    "context"
    "fmt"
)


type ListElectionsRequest struct {
    Context context.Context
}
func (o *ListElectionsRequest) validate() []string {
    ret := []string{}
    return ret
}


type CreateElectionRequest struct {
    Context context.Context
    Body Election
}
func (o *CreateElectionRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.Body.validate("body")...)
    return ret
}


type GetElectionRequest struct {
    Context context.Context
    PathParams GetElectionRequestPathParams
}
func (o *GetElectionRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.PathParams.validate()...)
    return ret
}

type GetElectionRequestPathParams struct {
    Id string `json:"id"`
}
func (o *GetElectionRequestPathParams) validate() []string {
     errors := []string{}
    idErrors := func(propValue string, parentName string) []string {
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
}(o.Id, "path")
    errors = append(errors, idErrors...)

    return errors
}

type UpdateElectionRequest struct {
    Context context.Context
    PathParams UpdateElectionRequestPathParams
    Body Election
}
func (o *UpdateElectionRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.PathParams.validate()...)
    ret = append(ret, o.Body.validate("body")...)
    return ret
}

type UpdateElectionRequestPathParams struct {
    Id string `json:"id"`
}
func (o *UpdateElectionRequestPathParams) validate() []string {
     errors := []string{}
    idErrors := func(propValue string, parentName string) []string {
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
}(o.Id, "path")
    errors = append(errors, idErrors...)

    return errors
}

type DeleteElectionRequest struct {
    Context context.Context
    PathParams DeleteElectionRequestPathParams
}
func (o *DeleteElectionRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.PathParams.validate()...)
    return ret
}

type DeleteElectionRequestPathParams struct {
    Id string `json:"id"`
}
func (o *DeleteElectionRequestPathParams) validate() []string {
     errors := []string{}
    idErrors := func(propValue string, parentName string) []string {
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
}(o.Id, "path")
    errors = append(errors, idErrors...)

    return errors
}

type StartElectionRequest struct {
    Context context.Context
    PathParams StartElectionRequestPathParams
    Body interface{}
}
func (o *StartElectionRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.PathParams.validate()...)
    return ret
}

type StartElectionRequestPathParams struct {
    Id string `json:"id"`
}
func (o *StartElectionRequestPathParams) validate() []string {
     errors := []string{}
    idErrors := func(propValue string, parentName string) []string {
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
}(o.Id, "path")
    errors = append(errors, idErrors...)

    return errors
}

type StopElectionRequest struct {
    Context context.Context
    PathParams StopElectionRequestPathParams
    Body interface{}
}
func (o *StopElectionRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.PathParams.validate()...)
    return ret
}

type StopElectionRequestPathParams struct {
    Id string `json:"id"`
}
func (o *StopElectionRequestPathParams) validate() []string {
     errors := []string{}
    idErrors := func(propValue string, parentName string) []string {
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
}(o.Id, "path")
    errors = append(errors, idErrors...)

    return errors
}

type GetBallotRequest struct {
    Context context.Context
    PathParams GetBallotRequestPathParams
}
func (o *GetBallotRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.PathParams.validate()...)
    return ret
}

type GetBallotRequestPathParams struct {
    Id string `json:"id"`
}
func (o *GetBallotRequestPathParams) validate() []string {
     errors := []string{}
    idErrors := func(propValue string, parentName string) []string {
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
}(o.Id, "path")
    errors = append(errors, idErrors...)

    return errors
}

type UpsertBallotRequest struct {
    Context context.Context
    PathParams UpsertBallotRequestPathParams
    Body Ballot
}
func (o *UpsertBallotRequest) validate() []string {
    ret := []string{}
    ret = append(ret, o.PathParams.validate()...)
    ret = append(ret, o.Body.validate("body")...)
    return ret
}

type UpsertBallotRequestPathParams struct {
    Id string `json:"id"`
}
func (o *UpsertBallotRequestPathParams) validate() []string {
     errors := []string{}
    idErrors := func(propValue string, parentName string) []string {
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
}(o.Id, "path")
    errors = append(errors, idErrors...)

    return errors
}

