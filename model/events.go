package model

type Event struct {
	ClubName              string      `json:"clubName"`
	Name                  string      `json:"name"`
	ToDate                string      `json:"toDate"`
	FromDate              string      `json:"fromDate"`
	ToTime                string      `json:"toTime"`
	FromTime              string      `json:"fromTime"`
	Budget                string      `json:"budget"`
	Description           string      `json:"description"`
	Category              string      `json:"category"`
	Venue                 string      `json:"venue"`
	Attendance            string      `json:"attendance"`
	ExpectedParticipants  string      `json:"expectedParticipants"`
	FacultyCoordinator    Participant `json:"sacultyCoordinator"`
	StudentCoordinator    Participant `json:"studentCoordinator"`
	GuestDetails          Guest       `json:"guest"`
	PROrequest            string      `json:"PROrequest"`
	CampusEngineerRequest string      `json:"campusEngineerRequest"`
	Duration              string      `json:"duration"`
	MainSponsor           Participant `json:"mainSponsor"`
}

type Guest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Gender         string `json:"gender"`
	Stake          string `json:"stake"`
	LocationOfStay string `json:"locationOfStay"`
}
