package model

import (
	"log"
	"sync"

	events "../../../lib"
)

func CreateEvent(e events.Event) error {
	c := make(chan error)
	//go createParticipant(e, "StudentCoordinator", c)
	result, err := events.Session.Run(`CREATE (n:EVENT {name:$name, clubName:$clubName, toDate:$toDate, 
		fromDate: $fromDate, toTime:$toTime, fromTime:$fromTime, budget:$budget, 
		description:$description, category:$category, venue:$venue, attendance:$attendance, 
		expectedParticipants:$expectedParticipants, PROrequest:$PROrequest, 
		campusEngineerRequest:$campusEngineerRequest, duration:$duration}) 
		RETURN n.name`, map[string]interface{}{

		"name":                  e.Name,
		"clubName":              e.ClubName,
		"toDate":                e.ToDate,
		"fromDate":              e.FromDate,
		"toTime":                e.ToTime,
		"fromTime":              e.FromTime,
		"budget":                e.Budget,
		"description":           e.Description,
		"category":              e.Category,
		"venue":                 e.Venue,
		"PROrequest":            e.PROrequest,
		"campusEngineerRequest": e.CampusEngineerRequest,
		"duration":              e.Duration,
		"attendance":            e.Attendance,
		"expectedParticipants":  e.ExpectedParticipants,
	})
	if err != nil {
		return err
	}

	result.Next()
	log.Println(result.Record().GetByIndex(0).(string))

	if err = result.Err(); err != nil {
		return err
	}

	// CREATE STUDENT COORDINATOR, FACULTY COORDINATOR, SPONSOR AND GUEST NODES
	var mutex = &sync.Mutex{}
	go events.CreateParticipant(e, "StudentCoordinator", c, mutex)
	go events.CreateParticipant(e, "FacultyCoordinator", c, mutex)
	go events.CreateParticipant(e, "MainSponsor", c, mutex)
	go events.CreateGuest(e, c, mutex)

	err1, err2, err3, err4 := <-c, <-c, <-c, <-c

	switch {
	case err1 != nil:
		return err1
	case err2 != nil:
		return err2
	case err3 != nil:
		return err3
	case err4 != nil:
		return err4
	}

	log.Println("Created Event node")
	return nil
}
