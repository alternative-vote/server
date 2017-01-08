package generated

import "io"

var _, _ = io.Pipe()

type ListElectionsResponse struct {
    StatusCode int
    Headers map[string]string
    Body []Election
}

type CreateElectionResponse struct {
    StatusCode int
    Headers map[string]string
    Body Election
    }

type GetElectionResponse struct {
    StatusCode int
    Headers map[string]string
    Body Election
    }

type UpdateElectionResponse struct {
    StatusCode int
    Headers map[string]string
    Body Election
    }

type DeleteElectionResponse struct {
    StatusCode int
    Headers map[string]string
    Body interface{}
    }

type StartElectionResponse struct {
    StatusCode int
    Headers map[string]string
    Body Election
    }

type StopElectionResponse struct {
    StatusCode int
    Headers map[string]string
    Body Election
    }

type LoginResponse struct {
    StatusCode int
    Headers map[string]string
    Body interface{}
    }

type GetBallotResponse struct {
    StatusCode int
    Headers map[string]string
    Body GetBallotResponseBody
    }

type UpdateBallotResponse struct {
    StatusCode int
    Headers map[string]string
    Body interface{}
    }

