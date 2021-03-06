/*
This file was generated.  Don't bother trying to edit anything here.
*/


package generated

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"fmt"

	"github.com/gorilla/mux"
)

var _ = ioutil.Discard
var _ = strconv.IntSize
var _, _ = io.Pipe()
var _ = fmt.Sprint("boggle")
var _ = strings.Compare("a", "b")
var _ = json.Decoder{}

var RouterElectionController IElectionController
var RouterAuthenticationController IAuthenticationController

func Router() *mux.Router {
    router := mux.NewRouter()
	router.KeepContext = true

    //get /elections/
    router.HandleFunc("/elections/", func(res http.ResponseWriter, req *http.Request) {
       request := new(ListElectionsRequest)
        request.Context = req.Context()
        errors := []string{}
    
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.ListElections(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("get")    
    //post /elections/
    router.HandleFunc("/elections/", func(res http.ResponseWriter, req *http.Request) {
       request := new(CreateElectionRequest)
        request.Context = req.Context()
        errors := []string{}
    
        bodyBytes, readErr := ioutil.ReadAll(req.Body)
        if readErr != nil {
            panic("error reading body from the request: " + readErr.Error())
        }
        
        //ok, now let's decode it into the actual request object
        bodyError := json.Unmarshal(bodyBytes, &request.Body)
        if bodyError != nil {
        	errors = append(errors, "invalid JSON or wrong types: "+bodyError.Error())
        }
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.CreateElection(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("post")    
    //options /elections/
	router.HandleFunc("/elections/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {
	        "/elections/{id}": {
	            "GET": "getElection",
	            "PUT": "updateElection",
	            "DELETE": "deleteElection"
	        }
	    },
	    "methods": {
	        "get": {
	            "operationId": "listElections",
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "array",
	                        "items": {
	                            "type": "object",
	                            "properties": {
	                                "id": {
	                                    "type": "string",
	                                    "format": "uuid"
	                                },
	                                "dateCreated": {
	                                    "type": "string",
	                                    "format": "date-time"
	                                },
	                                "dateUpdated": {
	                                    "type": "string",
	                                    "format": "date-time"
	                                },
	                                "title": {
	                                    "type": "string"
	                                },
	                                "subtitle": {
	                                    "type": "string"
	                                },
	                                "description": {
	                                    "type": "string"
	                                },
	                                "startDate": {
	                                    "type": "object",
	                                    "properties": {
	                                        "manual": {
	                                            "type": "boolean",
	                                            "default": true
	                                        },
	                                        "date": {
	                                            "type": "string",
	                                            "format": "date-time"
	                                        }
	                                    }
	                                },
	                                "endDate": {
	                                    "type": "object",
	                                    "properties": {
	                                        "manual": {
	                                            "type": "boolean",
	                                            "default": true
	                                        },
	                                        "date": {
	                                            "type": "string",
	                                            "format": "date-time"
	                                        }
	                                    }
	                                },
	                                "dateStarted": {
	                                    "type": "string",
	                                    "format": "date-time"
	                                },
	                                "dateEnded": {
	                                    "type": "string",
	                                    "format": "date-time"
	                                },
	                                "state": {
	                                    "type": "string"
	                                },
	                                "owner": {
	                                    "type": "object",
	                                    "properties": {
	                                        "id": {
	                                            "type": "string",
	                                            "format": "uuid"
	                                        },
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "isAccount": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                },
	                                "voters": {
	                                    "type": "array",
	                                    "items": {
	                                        "type": "object",
	                                        "properties": {
	                                            "email": {
	                                                "type": "string",
	                                                "format": "email"
	                                            },
	                                            "voteEmailSent": {
	                                                "type": "boolean"
	                                            },
	                                            "resultsEmailSent": {
	                                                "type": "boolean"
	                                            }
	                                        }
	                                    }
	                                },
	                                "candidates": {
	                                    "type": "array",
	                                    "items": {
	                                        "type": "object",
	                                        "properties": {
	                                            "title": {
	                                                "type": "string"
	                                            },
	                                            "subtitle": {
	                                                "type": "string"
	                                            },
	                                            "description": {
	                                                "type": "string"
	                                            }
	                                        }
	                                    }
	                                },
	                                "results": {
	                                    "type": "object",
	                                    "properties": {
	                                        "orderedCandidates": {
	                                            "type": "array",
	                                            "items": {
	                                                "type": "object",
	                                                "properties": {
	                                                    "title": {
	                                                        "type": "string"
	                                                    },
	                                                    "subtitle": {
	                                                        "type": "string"
	                                                    },
	                                                    "description": {
	                                                        "type": "string"
	                                                    }
	                                                }
	                                            }
	                                        },
	                                        "stats": {
	                                            "type": "object",
	                                            "properties": {
	                                                "start": {
	                                                    "type": "string",
	                                                    "format": "date-time"
	                                                },
	                                                "end": {
	                                                    "type": "string",
	                                                    "format": "date-time"
	                                                },
	                                                "numVoters": {
	                                                    "type": "integer"
	                                                },
	                                                "ballotsSubmitted": {
	                                                    "type": "integer"
	                                                },
	                                                "averageCandidatesRanked": {
	                                                    "type": "number"
	                                                }
	                                            }
	                                        },
	                                        "fullData": {
	                                            "type": "array",
	                                            "items": {
	                                                "type": "object"
	                                            }
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        },
	        "post": {
	            "operationId": "createElection",
	            "security": [
	                {
	                    "Bearer": []
	                }
	            ],
	            "parameters": [
	                {
	                    "name": "body",
	                    "in": "body",
	                    "required": true,
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            ],
	            "responses": {
	                "201": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        }
	    }
	}`))
	}).Methods("options")
    
    //405 handler for /elections/
    router.HandleFunc("/elections/", func(res http.ResponseWriter, req *http.Request) {
	    panic(HttpError(405))
	})

    //get /elections/{id}/
    router.HandleFunc("/elections/{id}/", func(res http.ResponseWriter, req *http.Request) {
       request := new(GetElectionRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'id' in form data
        request.PathParams.Id = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "id is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("id: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["id"])
        
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.GetElection(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("get")    
    //put /elections/{id}/
    router.HandleFunc("/elections/{id}/", func(res http.ResponseWriter, req *http.Request) {
       request := new(UpdateElectionRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'id' in form data
        request.PathParams.Id = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "id is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("id: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["id"])
        
        bodyBytes, readErr := ioutil.ReadAll(req.Body)
        if readErr != nil {
            panic("error reading body from the request: " + readErr.Error())
        }
        
        //ok, now let's decode it into the actual request object
        bodyError := json.Unmarshal(bodyBytes, &request.Body)
        if bodyError != nil {
        	errors = append(errors, "invalid JSON or wrong types: "+bodyError.Error())
        }
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.UpdateElection(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("put")    
    //delete /elections/{id}/
    router.HandleFunc("/elections/{id}/", func(res http.ResponseWriter, req *http.Request) {
       request := new(DeleteElectionRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'id' in form data
        request.PathParams.Id = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "id is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("id: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["id"])
        
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.DeleteElection(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("delete")    
    //options /elections/{id}/
	router.HandleFunc("/elections/{id}/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {
	        "/elections/{id}/actions/start": {
	            "PUT": "startElection"
	        },
	        "/elections/{id}/actions/stop": {
	            "PUT": "stopElection"
	        },
	        "/elections/{id}/actions/sendEmail": {
	            "PUT": "sendEmail"
	        }
	    },
	    "methods": {
	        "get": {
	            "operationId": "getElection",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "id",
	                    "required": true,
	                    "type": "string",
	                    "format": "uuid"
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        },
	        "put": {
	            "operationId": "updateElection",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "id",
	                    "required": true,
	                    "type": "string",
	                    "format": "uuid"
	                },
	                {
	                    "name": "body",
	                    "in": "body",
	                    "required": true,
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        },
	        "delete": {
	            "operationId": "deleteElection",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "id",
	                    "required": true,
	                    "type": "string",
	                    "format": "uuid"
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object"
	                    }
	                }
	            }
	        }
	    }
	}`))
	}).Methods("options")
    
    //405 handler for /elections/{id}/
    router.HandleFunc("/elections/{id}/", func(res http.ResponseWriter, req *http.Request) {
	    panic(HttpError(405))
	})

    //put /elections/{id}/actions/start/
    router.HandleFunc("/elections/{id}/actions/start/", func(res http.ResponseWriter, req *http.Request) {
       request := new(StartElectionRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'id' in form data
        request.PathParams.Id = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "id is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("id: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["id"])
        
        bodyBytes, readErr := ioutil.ReadAll(req.Body)
        if readErr != nil {
            panic("error reading body from the request: " + readErr.Error())
        }
        
        //ok, now let's decode it into the actual request object
        bodyError := json.Unmarshal(bodyBytes, &request.Body)
        if bodyError != nil {
        	errors = append(errors, "invalid JSON or wrong types: "+bodyError.Error())
        }
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.StartElection(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("put")    
    //options /elections/{id}/actions/start/
	router.HandleFunc("/elections/{id}/actions/start/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {},
	    "methods": {
	        "put": {
	            "operationId": "startElection",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "id",
	                    "required": true,
	                    "type": "string",
	                    "format": "uuid"
	                },
	                {
	                    "name": "body",
	                    "in": "body",
	                    "required": true,
	                    "schema": {
	                        "type": "object"
	                    }
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        }
	    }
	}`))
	}).Methods("options")
    
    //405 handler for /elections/{id}/actions/start/
    router.HandleFunc("/elections/{id}/actions/start/", func(res http.ResponseWriter, req *http.Request) {
	    panic(HttpError(405))
	})

    //put /elections/{id}/actions/stop/
    router.HandleFunc("/elections/{id}/actions/stop/", func(res http.ResponseWriter, req *http.Request) {
       request := new(StopElectionRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'id' in form data
        request.PathParams.Id = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "id is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("id: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["id"])
        
        bodyBytes, readErr := ioutil.ReadAll(req.Body)
        if readErr != nil {
            panic("error reading body from the request: " + readErr.Error())
        }
        
        //ok, now let's decode it into the actual request object
        bodyError := json.Unmarshal(bodyBytes, &request.Body)
        if bodyError != nil {
        	errors = append(errors, "invalid JSON or wrong types: "+bodyError.Error())
        }
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.StopElection(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("put")    
    //options /elections/{id}/actions/stop/
	router.HandleFunc("/elections/{id}/actions/stop/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {},
	    "methods": {
	        "put": {
	            "operationId": "stopElection",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "id",
	                    "required": true,
	                    "type": "string",
	                    "format": "uuid"
	                },
	                {
	                    "name": "body",
	                    "in": "body",
	                    "required": true,
	                    "schema": {
	                        "type": "object"
	                    }
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        }
	    }
	}`))
	}).Methods("options")
    
    //405 handler for /elections/{id}/actions/stop/
    router.HandleFunc("/elections/{id}/actions/stop/", func(res http.ResponseWriter, req *http.Request) {
	    panic(HttpError(405))
	})

    //put /elections/{id}/actions/sendEmail/
    router.HandleFunc("/elections/{id}/actions/sendEmail/", func(res http.ResponseWriter, req *http.Request) {
       request := new(SendEmailRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'id' in form data
        request.PathParams.Id = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "id is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("id: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["id"])
        
        //'force' in form data
        request.QueryParams.Force = func(s string) string {
            var ret string
            if s == "" {
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("force: '%v' is not a valid string", s))
            }
        
            return ret
        }(req.URL.Query().Get("force"))
        
        bodyBytes, readErr := ioutil.ReadAll(req.Body)
        if readErr != nil {
            panic("error reading body from the request: " + readErr.Error())
        }
        
        //ok, now let's decode it into the actual request object
        bodyError := json.Unmarshal(bodyBytes, &request.Body)
        if bodyError != nil {
        	errors = append(errors, "invalid JSON or wrong types: "+bodyError.Error())
        }
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.SendEmail(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("put")    
    //options /elections/{id}/actions/sendEmail/
	router.HandleFunc("/elections/{id}/actions/sendEmail/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {},
	    "methods": {
	        "put": {
	            "operationId": "sendEmail",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "id",
	                    "required": true,
	                    "type": "string",
	                    "format": "uuid"
	                },
	                {
	                    "name": "force",
	                    "in": "query",
	                    "required": false,
	                    "type": "string"
	                },
	                {
	                    "name": "body",
	                    "in": "body",
	                    "required": true,
	                    "schema": {
	                        "type": "object"
	                    }
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "dateCreated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateUpdated": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "title": {
	                                "type": "string"
	                            },
	                            "subtitle": {
	                                "type": "string"
	                            },
	                            "description": {
	                                "type": "string"
	                            },
	                            "startDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "endDate": {
	                                "type": "object",
	                                "properties": {
	                                    "manual": {
	                                        "type": "boolean",
	                                        "default": true
	                                    },
	                                    "date": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    }
	                                }
	                            },
	                            "dateStarted": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "dateEnded": {
	                                "type": "string",
	                                "format": "date-time"
	                            },
	                            "state": {
	                                "type": "string"
	                            },
	                            "owner": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "email": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "isAccount": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            },
	                            "voters": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "email": {
	                                            "type": "string",
	                                            "format": "email"
	                                        },
	                                        "voteEmailSent": {
	                                            "type": "boolean"
	                                        },
	                                        "resultsEmailSent": {
	                                            "type": "boolean"
	                                        }
	                                    }
	                                }
	                            },
	                            "candidates": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "results": {
	                                "type": "object",
	                                "properties": {
	                                    "orderedCandidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "stats": {
	                                        "type": "object",
	                                        "properties": {
	                                            "start": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "end": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            },
	                                            "numVoters": {
	                                                "type": "integer"
	                                            },
	                                            "ballotsSubmitted": {
	                                                "type": "integer"
	                                            },
	                                            "averageCandidatesRanked": {
	                                                "type": "number"
	                                            }
	                                        }
	                                    },
	                                    "fullData": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object"
	                                        }
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        }
	    }
	}`))
	}).Methods("options")
    
    //405 handler for /elections/{id}/actions/sendEmail/
    router.HandleFunc("/elections/{id}/actions/sendEmail/", func(res http.ResponseWriter, req *http.Request) {
	    panic(HttpError(405))
	})

    //post /login/
    router.HandleFunc("/login/", func(res http.ResponseWriter, req *http.Request) {
       request := new(LoginRequest)
        request.Context = req.Context()
        errors := []string{}
    
        bodyBytes, readErr := ioutil.ReadAll(req.Body)
        if readErr != nil {
            panic("error reading body from the request: " + readErr.Error())
        }
        
        //ok, now let's decode it into the actual request object
        bodyError := json.Unmarshal(bodyBytes, &request.Body)
        if bodyError != nil {
        	errors = append(errors, "invalid JSON or wrong types: "+bodyError.Error())
        }
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterAuthenticationController == nil {
            panic(HttpError(501))
        }
        response := RouterAuthenticationController.Login(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("post")    
    //options /login/
	router.HandleFunc("/login/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {},
	    "methods": {
	        "post": {
	            "operationId": "login",
	            "parameters": [
	                {
	                    "name": "body",
	                    "in": "body",
	                    "required": true,
	                    "schema": {
	                        "type": "object",
	                        "required": [
	                            "email",
	                            "password"
	                        ],
	                        "properties": {
	                            "email": {
	                                "type": "string",
	                                "format": "email"
	                            },
	                            "password": {
	                                "type": "string"
	                            }
	                        }
	                    }
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object"
	                    }
	                }
	            }
	        }
	    }
	}`))
	}).Methods("options")
    
    //405 handler for /login/
    router.HandleFunc("/login/", func(res http.ResponseWriter, req *http.Request) {
	    panic(HttpError(405))
	})

    //get /vote/{token}/
    router.HandleFunc("/vote/{token}/", func(res http.ResponseWriter, req *http.Request) {
       request := new(GetBallotRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'token' in form data
        request.PathParams.Token = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "token is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("token: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["token"])
        
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.GetBallot(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("get")    
    //put /vote/{token}/
    router.HandleFunc("/vote/{token}/", func(res http.ResponseWriter, req *http.Request) {
       request := new(UpdateBallotRequest)
        request.Context = req.Context()
        errors := []string{}
    
        //'token' in form data
        request.PathParams.Token = func(s string) string {
            var ret string
            if s == "" {
                errors = append(errors, "token is a required path parameter")
                return ret
            }
            var err error
            ret, err = func(s string) (string, error) {
            return s, nil
        }(s)
            if err != nil {
                errors = append(errors, fmt.Sprintf("token: '%v' is not a valid string", s))
            }
        
            return ret
        }(mux.Vars(req)["token"])
        
        bodyBytes, readErr := ioutil.ReadAll(req.Body)
        if readErr != nil {
            panic("error reading body from the request: " + readErr.Error())
        }
        
        //ok, now let's decode it into the actual request object
        bodyError := json.Unmarshal(bodyBytes, &request.Body)
        if bodyError != nil {
        	errors = append(errors, "invalid JSON or wrong types: "+bodyError.Error())
        }
        
        //don't bother to validate if we had deserialization errors
        if len(errors) == 0 {
            errors = append(errors, request.validate()...)
        }
        
    
        if len(errors) > 0 {
            apiError := HttpError(400)
            apiError.ValidationErrors = errors
            panic(apiError)
    	}
        if RouterElectionController == nil {
            panic(HttpError(501))
        }
        response := RouterElectionController.UpdateBallot(request)
    
        //transfer headers to the actual http response
        if response != nil {
            for k, v := range response.Headers {
                res.Header().Set(k, v)
            }
        }
        
        //transfer status code to the actual http response
        //order matters - make sure to do this after setting other header values!
        res.WriteHeader(response.StatusCode)
        
        responseBytes, _ := json.Marshal(response.Body)
        res.Write(responseBytes)
    }).Methods("put")    
    //options /vote/{token}/
	router.HandleFunc("/vote/{token}/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {},
	    "methods": {
	        "get": {
	            "operationId": "getBallot",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "token",
	                    "required": true,
	                    "type": "string"
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "election": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "dateCreated": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "dateUpdated": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "title": {
	                                        "type": "string"
	                                    },
	                                    "subtitle": {
	                                        "type": "string"
	                                    },
	                                    "description": {
	                                        "type": "string"
	                                    },
	                                    "startDate": {
	                                        "type": "object",
	                                        "properties": {
	                                            "manual": {
	                                                "type": "boolean",
	                                                "default": true
	                                            },
	                                            "date": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            }
	                                        }
	                                    },
	                                    "endDate": {
	                                        "type": "object",
	                                        "properties": {
	                                            "manual": {
	                                                "type": "boolean",
	                                                "default": true
	                                            },
	                                            "date": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            }
	                                        }
	                                    },
	                                    "dateStarted": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "dateEnded": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "state": {
	                                        "type": "string"
	                                    },
	                                    "owner": {
	                                        "type": "object",
	                                        "properties": {
	                                            "id": {
	                                                "type": "string",
	                                                "format": "uuid"
	                                            },
	                                            "email": {
	                                                "type": "string",
	                                                "format": "email"
	                                            },
	                                            "isAccount": {
	                                                "type": "boolean"
	                                            }
	                                        }
	                                    },
	                                    "voters": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "email": {
	                                                    "type": "string",
	                                                    "format": "email"
	                                                },
	                                                "voteEmailSent": {
	                                                    "type": "boolean"
	                                                },
	                                                "resultsEmailSent": {
	                                                    "type": "boolean"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "candidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "results": {
	                                        "type": "object",
	                                        "properties": {
	                                            "orderedCandidates": {
	                                                "type": "array",
	                                                "items": {
	                                                    "type": "object",
	                                                    "properties": {
	                                                        "title": {
	                                                            "type": "string"
	                                                        },
	                                                        "subtitle": {
	                                                            "type": "string"
	                                                        },
	                                                        "description": {
	                                                            "type": "string"
	                                                        }
	                                                    }
	                                                }
	                                            },
	                                            "stats": {
	                                                "type": "object",
	                                                "properties": {
	                                                    "start": {
	                                                        "type": "string",
	                                                        "format": "date-time"
	                                                    },
	                                                    "end": {
	                                                        "type": "string",
	                                                        "format": "date-time"
	                                                    },
	                                                    "numVoters": {
	                                                        "type": "integer"
	                                                    },
	                                                    "ballotsSubmitted": {
	                                                        "type": "integer"
	                                                    },
	                                                    "averageCandidatesRanked": {
	                                                        "type": "number"
	                                                    }
	                                                }
	                                            },
	                                            "fullData": {
	                                                "type": "array",
	                                                "items": {
	                                                    "type": "object"
	                                                }
	                                            }
	                                        }
	                                    }
	                                }
	                            },
	                            "ballot": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string"
	                                    },
	                                    "electionId": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "voter": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "votes": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "isSubmitted": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        },
	        "put": {
	            "operationId": "updateBallot",
	            "parameters": [
	                {
	                    "in": "path",
	                    "name": "token",
	                    "required": true,
	                    "type": "string"
	                },
	                {
	                    "name": "body",
	                    "in": "body",
	                    "required": true,
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "id": {
	                                "type": "string"
	                            },
	                            "electionId": {
	                                "type": "string",
	                                "format": "uuid"
	                            },
	                            "voter": {
	                                "type": "string",
	                                "format": "email"
	                            },
	                            "votes": {
	                                "type": "array",
	                                "items": {
	                                    "type": "object",
	                                    "properties": {
	                                        "title": {
	                                            "type": "string"
	                                        },
	                                        "subtitle": {
	                                            "type": "string"
	                                        },
	                                        "description": {
	                                            "type": "string"
	                                        }
	                                    }
	                                }
	                            },
	                            "isSubmitted": {
	                                "type": "boolean"
	                            }
	                        }
	                    }
	                }
	            ],
	            "responses": {
	                "200": {
	                    "description": "something",
	                    "schema": {
	                        "type": "object",
	                        "properties": {
	                            "election": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "dateCreated": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "dateUpdated": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "title": {
	                                        "type": "string"
	                                    },
	                                    "subtitle": {
	                                        "type": "string"
	                                    },
	                                    "description": {
	                                        "type": "string"
	                                    },
	                                    "startDate": {
	                                        "type": "object",
	                                        "properties": {
	                                            "manual": {
	                                                "type": "boolean",
	                                                "default": true
	                                            },
	                                            "date": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            }
	                                        }
	                                    },
	                                    "endDate": {
	                                        "type": "object",
	                                        "properties": {
	                                            "manual": {
	                                                "type": "boolean",
	                                                "default": true
	                                            },
	                                            "date": {
	                                                "type": "string",
	                                                "format": "date-time"
	                                            }
	                                        }
	                                    },
	                                    "dateStarted": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "dateEnded": {
	                                        "type": "string",
	                                        "format": "date-time"
	                                    },
	                                    "state": {
	                                        "type": "string"
	                                    },
	                                    "owner": {
	                                        "type": "object",
	                                        "properties": {
	                                            "id": {
	                                                "type": "string",
	                                                "format": "uuid"
	                                            },
	                                            "email": {
	                                                "type": "string",
	                                                "format": "email"
	                                            },
	                                            "isAccount": {
	                                                "type": "boolean"
	                                            }
	                                        }
	                                    },
	                                    "voters": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "email": {
	                                                    "type": "string",
	                                                    "format": "email"
	                                                },
	                                                "voteEmailSent": {
	                                                    "type": "boolean"
	                                                },
	                                                "resultsEmailSent": {
	                                                    "type": "boolean"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "candidates": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "results": {
	                                        "type": "object",
	                                        "properties": {
	                                            "orderedCandidates": {
	                                                "type": "array",
	                                                "items": {
	                                                    "type": "object",
	                                                    "properties": {
	                                                        "title": {
	                                                            "type": "string"
	                                                        },
	                                                        "subtitle": {
	                                                            "type": "string"
	                                                        },
	                                                        "description": {
	                                                            "type": "string"
	                                                        }
	                                                    }
	                                                }
	                                            },
	                                            "stats": {
	                                                "type": "object",
	                                                "properties": {
	                                                    "start": {
	                                                        "type": "string",
	                                                        "format": "date-time"
	                                                    },
	                                                    "end": {
	                                                        "type": "string",
	                                                        "format": "date-time"
	                                                    },
	                                                    "numVoters": {
	                                                        "type": "integer"
	                                                    },
	                                                    "ballotsSubmitted": {
	                                                        "type": "integer"
	                                                    },
	                                                    "averageCandidatesRanked": {
	                                                        "type": "number"
	                                                    }
	                                                }
	                                            },
	                                            "fullData": {
	                                                "type": "array",
	                                                "items": {
	                                                    "type": "object"
	                                                }
	                                            }
	                                        }
	                                    }
	                                }
	                            },
	                            "ballot": {
	                                "type": "object",
	                                "properties": {
	                                    "id": {
	                                        "type": "string"
	                                    },
	                                    "electionId": {
	                                        "type": "string",
	                                        "format": "uuid"
	                                    },
	                                    "voter": {
	                                        "type": "string",
	                                        "format": "email"
	                                    },
	                                    "votes": {
	                                        "type": "array",
	                                        "items": {
	                                            "type": "object",
	                                            "properties": {
	                                                "title": {
	                                                    "type": "string"
	                                                },
	                                                "subtitle": {
	                                                    "type": "string"
	                                                },
	                                                "description": {
	                                                    "type": "string"
	                                                }
	                                            }
	                                        }
	                                    },
	                                    "isSubmitted": {
	                                        "type": "boolean"
	                                    }
	                                }
	                            }
	                        }
	                    }
	                }
	            }
	        }
	    }
	}`))
	}).Methods("options")
    
    //405 handler for /vote/{token}/
    router.HandleFunc("/vote/{token}/", func(res http.ResponseWriter, req *http.Request) {
	    panic(HttpError(405))
	})


	//root options request handler
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	    res.Write([]byte(`{
	    "childRoutes": {
	        "/elections": {
	            "GET": "listElections",
	            "POST": "createElection"
	        },
	        "/login": {
	            "POST": "login"
	        },
	        "/vote/{token}": {
	            "GET": "getBallot",
	            "PUT": "updateBallot"
	        }
	    },
	    "methods": {}
	}`))
	}).Methods("options")

    //catch-all 404 handler
    router.HandleFunc("/{restOfRoute:.*}", func(res http.ResponseWriter, req *http.Request) {
		panic(HttpError(404))
	})

	return router

}

