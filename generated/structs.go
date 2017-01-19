package generated

type ElectionResultsStats struct {
    Start APITime `json:"start"`
    End APITime `json:"end"`
    NumVoters int64 `json:"numVoters"`
    BallotsSubmitted int64 `json:"ballotsSubmitted"`
    AverageCandidatesRanked float64 `json:"averageCandidatesRanked"`
    MetaData SerializationMetadata `json:"-"`
}

type ElectionResults struct {
    OrderedCandidates []Candidate `json:"orderedCandidates"`
    Stats ElectionResultsStats `json:"stats"`
    FullData []interface{} `json:"fullData"`
    MetaData SerializationMetadata `json:"-"`
}

type Election struct {
    Id string `json:"id"`
    DateCreated APITime `json:"dateCreated"`
    DateUpdated APITime `json:"dateUpdated"`
    Title string `json:"title"`
    Subtitle string `json:"subtitle"`
    Description string `json:"description"`
    StartDate Timer `json:"startDate"`
    EndDate Timer `json:"endDate"`
    DateStarted APITime `json:"dateStarted"`
    DateEnded APITime `json:"dateEnded"`
    State string `json:"state"`
    Owner User `json:"owner"`
    Voters []Voter `json:"voters"`
    Candidates []Candidate `json:"candidates"`
    Results ElectionResults `json:"results"`
    MetaData SerializationMetadata `json:"-"`
}

type Voter struct {
    Email string `json:"email"`
    VoteEmailSent bool `json:"voteEmailSent"`
    ResultsEmailSent bool `json:"resultsEmailSent"`
    MetaData SerializationMetadata `json:"-"`
}

type Timer struct {
    Manual bool `json:"manual"`
    Date APITime `json:"date"`
    MetaData SerializationMetadata `json:"-"`
}

type Ballot struct {
    Id string `json:"id"`
    ElectionId string `json:"electionId"`
    Voter string `json:"voter"`
    Votes []Candidate `json:"votes"`
    IsSubmitted bool `json:"isSubmitted"`
    MetaData SerializationMetadata `json:"-"`
}

type User struct {
    Id string `json:"id"`
    Email string `json:"email"`
    IsAccount bool `json:"isAccount"`
    MetaData SerializationMetadata `json:"-"`
}

type Candidate struct {
    Title string `json:"title"`
    Subtitle string `json:"subtitle"`
    Description string `json:"description"`
    MetaData SerializationMetadata `json:"-"`
}

type LoginRequestBody struct {
    Email string `json:"email"`
    Password string `json:"password"`
    MetaData SerializationMetadata `json:"-"`
}

type GetBallotResponseBody struct {
    Election Election `json:"election"`
    Ballot Ballot `json:"ballot"`
    MetaData SerializationMetadata `json:"-"`
}

type UpdateBallotResponseBody struct {
    Election Election `json:"election"`
    Ballot Ballot `json:"ballot"`
    MetaData SerializationMetadata `json:"-"`
}

