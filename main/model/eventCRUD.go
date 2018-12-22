package model

import (
	"log"
	"sync"
)

func CreateEvent(e Event) error {
	c := make(chan error)
	//go createParticipant(e, "StudentCoordinator", c)
	result, err := session.Run(`CREATE (n:EVENT {name:$name, clubName:$clubName, ToDate:$toDate, 
		FromDate: $fromDate, ToTime:$toTime, FromTime:$fromTime, Budget:$budget, 
		Description:$description, Category:$category, Venue:$venue, Attendance:$attendance, 
		ExpectedParticipants:$expectedParticipants, PROrequest:$PROrequest, 
		CampusEngineerRequest:$campusEngineerRequest, Duration:$duration}) 
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
	go createParticipant(e, "StudentCoordinator", c, mutex)
	go createParticipant(e, "FacultyCoordinator", c, mutex)
	go createParticipant(e, "MainSponsor", c, mutex)
	go createGuest(e, c, mutex)

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

// create a new node with given label and participant data struct
func createParticipant(e Event, label string, c chan error, mutex *sync.Mutex) {
	mutex.Lock()
	result, err := session.Run(`MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:INCHARGE {name:$name, registrationNumber:$registrationNumber,
		email:$email, phoneNumber:$phoneNumber, gender: $gender})<-[:`+label+`]-(a) `, map[string]interface{}{
		"EventName":          e.Name,
		"name":               e.getField(label, "Name"),
		"registrationNumber": e.getField(label, "RegistrationNumber"),
		"email":              e.getField(label, "Email"),
		"phoneNumber":        e.getField(label, "PhoneNumber"),
		"gender":             e.getField(label, "Gender"),
	})
	if err != nil {
		c <- err
	}
	mutex.Unlock()

	if err = result.Err(); err != nil {
		c <- err
	}
	log.Printf("Created %s node", label)
	c <- nil
}

// create a new guest node with relationship to the event
func createGuest(e Event, c chan error, mutex *sync.Mutex) {
	mutex.Lock()
	result, err := session.Run(`MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:GUEST {name:$name, stake:$stake,
	email:$email, phoneNumber:$phoneNumber, gender: $gender, locationOfStay:$locationOfStay
	})<-[:GUEST]-(a) `, map[string]interface{}{
		"EventName":      e.Name,
		"name":           e.getField("GuestDetails", "Name"),
		"stake":          e.getField("GuestDetails", "Stake"),
		"email":          e.getField("GuestDetails", "Email"),
		"phoneNumber":    e.getField("GuestDetails", "PhoneNumber"),
		"gender":         e.getField("GuestDetails", "Gender"),
		"locationOfStay": e.getField("GuestDetails", "LocationOfStay"),
	})
	if err != nil {
		c <- err
	}
	mutex.Unlock()

	if err = result.Err(); err != nil {
		c <- err
	}
	log.Println("Created GUEST node")
	c <- nil
}
