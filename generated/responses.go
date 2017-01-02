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

type GetBallotResponse struct {
    StatusCode int
    Headers map[string]string
    Body Ballot
    }

type UpsertBallotResponse struct {
    StatusCode int
    Headers map[string]string
    Body Ballot
    }

type LoginResponse struct {
    StatusCode int
    Headers map[string]string
    Body interface{}
    }

