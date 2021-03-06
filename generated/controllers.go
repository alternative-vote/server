package generated

import "net/http"

var _ = http.MethodGet

type IElectionController interface{
    ListElections(*ListElectionsRequest) *ListElectionsResponse
    
    CreateElection(*CreateElectionRequest) *CreateElectionResponse
    
    GetElection(*GetElectionRequest) *GetElectionResponse
    
    UpdateElection(*UpdateElectionRequest) *UpdateElectionResponse
    
    DeleteElection(*DeleteElectionRequest) *DeleteElectionResponse
    
    StartElection(*StartElectionRequest) *StartElectionResponse
    
    StopElection(*StopElectionRequest) *StopElectionResponse
    
    SendEmail(*SendEmailRequest) *SendEmailResponse
    
    GetBallot(*GetBallotRequest) *GetBallotResponse
    
    UpdateBallot(*UpdateBallotRequest) *UpdateBallotResponse
    
}

type IAuthenticationController interface{
    Login(*LoginRequest) *LoginResponse
    
}

