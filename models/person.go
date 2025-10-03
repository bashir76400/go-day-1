package models

type Person struct {
	Name     string
	Age      int
	IDNumber int
}

type Candidate struct {
	Name   string
	Party  string
	Votes  int
	Voters []Person
}
