package generated

type ElectionRoles struct {
    Voters Role `json:"voters"`
    Administrators Role `json:"administrators"`
    MetaData SerializationMetadata `json:"-"`
}

type ElectionResults struct {
    OrderedCandidates []string `json:"orderedCandidates"`
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
    Roles ElectionRoles `json:"roles"`
    Candidates []Candidate `json:"candidates"`
    Results ElectionResults `json:"results"`
    MetaData SerializationMetadata `json:"-"`
}

type Role struct {
    IsPublic bool `json:"isPublic"`
    Members []User `json:"members"`
    MetaData SerializationMetadata `json:"-"`
}

type Timer struct {
    Manual bool `json:"manual"`
    Date APITime `json:"date"`
    MetaData SerializationMetadata `json:"-"`
}

type Ballot struct {
    Voter User `json:"voter"`
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

