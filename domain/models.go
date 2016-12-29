package domain

import "github.com/alternative-vote/server/generated"

type Election struct {
	generated.Election
	Ballots []generated.Ballot
}

// type Election struct {
// 	ID          string
// 	DateCreated time.Time
// 	DateUpdated time.Time
// }

// type Timer struct {
// 	Manual bool
// 	Time   time.Timer
// }

// type User struct {
// 	ID        string
// 	Email     string
// 	Password  string
// 	IsAccount bool
// }

// type Role struct {
// 	IsPublic bool
// 	Members  []User
// }

// type Candidate struct {
// 	ID          string
// 	Title       string
// 	Subtitle    string
// 	Description string
// }

// type Ballot struct {
// }
