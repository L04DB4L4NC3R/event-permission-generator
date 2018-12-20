# Project omega
The last event coordinator you will ever need

<br />

## Main interface


* CRUD an event

* Check bookings

* Permission generator

* Auto update website


<br />



#### Event schema

<br />

```go
{
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
}

```

<br />

#### Guest schema

<br />

```go
{
	Name           string `json:"name"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phoneNumber"`
	Gender         string `json:"gender"`
	Stake          string `json:"stake"`
	LocationOfStay string `json:"locationOfStay"`
}
```


<br />

#### Participant schema

<br />

```go
{
	Name               string `json:"name"`
	RegistrationNumber string `json:"registrationNumber"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	Gender             string `json:"gender"`
	EventsAttended     string `json:"eventsAttended"`
}
```

<br />

#### Form generate request

<br />

```json
{
   "clubName":"GDG",
   "name":"DEVFEST",
   "toDate":"10TH OCTOBER",
   "fromDate":"8TH OCTOBER",
   "toTime":"10 PM",
   "fromTime":"11 AM",
   "budget":"200000",
   "description":"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING",
   "category":"TECHNICAL",
   "venue":"ANNA AUDI",
   "attendance":"4000",
   "expectedParticipants":"4000",
   "facultyCoordinator":{
      "name":"NOORU MAA",
      "registrationNumber":"17BBE1010",
      "email":"SDADAS@A.COM",
      "phoneNumber":"919191991911",
      "gender":"M",
      "eventsAttended":"ALL"
   },
   "studentCoordinator":{
      "name":"NOORU BAAP",
      "registrationNumber":"17BBE1010",
      "email":"SDADAS@A.COM",
      "phoneNumber":"919191991911",
      "gender":"M",
      "eventsAttended":"ALL"
   },
   "guest":{
      "name":"ALLAHH DAAS",
      "email":"ASDSAD#ASD.COM",
      "phoneNumber":"11111111111",
      "gender":"F",
      "stake":"SOME MONAYYYY",
      "locationOfStay":"TERA GHAR"
   },
   "PROrequest":"SAJDOOSIJANDFSAKFDSAFD",
   "campusEngineerRequest":"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
   "duration":"16 hours",
   "mainSponsor":{
      "name":"ALLAHH DAAS",
      "email":"ASDSAD#ASD.COM",
      "phoneNumber":"11111111111",
      "gender":"F",
      "stake":"SOME MONAYYYY",
      "locationOfStay":"TERA GHAR"
   }
}
```