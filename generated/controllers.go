package generated

import "net/http"

var _ = http.MethodGet

type IElectionController interface{
    ListElections(*ListElectionsRequest) *ListElectionsResponse
    
    CreateElection(*CreateElectionRequest) *CreateElectionResponse
    
    GetElection(*GetElectionRequest) *GetElectionResponse
    
    UpdateElection(*UpdateElectionRequest) *UpdateElectionResponse
    
    StartElection(*StartElectionRequest) *StartElectionResponse
    
    StopElection(*StopElectionRequest) *StopElectionResponse
    
    GetBallot(*GetBallotRequest) *GetBallotResponse
    
    UpsertBallot(*UpsertBallotRequest) *UpsertBallotResponse
    
}

