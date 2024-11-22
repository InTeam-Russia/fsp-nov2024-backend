package event

import "time"

type Indicators struct {
	Dates []time.Time `json:"dates"`
}

type EventOut struct {
	EkpId             int64     `json:"ekpId"`
	Sport             string    `json:"sport"`
	Team              string    `json:"team"`
	CompetitionType   string    `json:"competitionType"`
	GenderAndAgeGroup []string  `json:"genderAndAgeGroup"`
	ProgramDiscipline []string  `json:"programDiscipline"`
	Venue             Venue     `json:"venue"`
	MembersCount      int       `json:"membersCount"`
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
}

type Venue struct {
	Country string  `json:"country"`
	Region  *string `json:"region,omitempty"`
	City    string  `json:"city"`
}

type EventFilter struct {
	Sport             *string
	Team              *string
	CompetitionType   *string
	GenderAndAgeGroup *string
	ProgramDiscipline *string
	Country           *string
	Region            *string
	City              *string
	MinMembersCount   *int
	MaxMembersCount   *int
	StartDate         *time.Time
	EndDate           *time.Time
}
