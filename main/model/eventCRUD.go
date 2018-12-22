package model

import "log"

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

	// CREATE STUDENT COORDINATOR AND FACULTY COORDINATOR NODES
	go createParticipant(e, "StudentCoordinator", c)
	go createParticipant(e, "FacultyCoordinator", c)
	if err1, err2 := <-c, <-c; err1 != nil || err2 != nil {
		if err1 != nil {
			return err1
		} else {
			return err2
		}
	}
	log.Println("Created Event node")
	return nil
}

// create a new node with given label and participant data struct
func createParticipant(e Event, label string, c chan error) {

	result, err := session.Run(`CREATE (n:`+label+` {name:$name, registrationNumber:$registrationNumber,
		email:$email, phoneNumber:$phoneNumber, gender: $gender}) `, map[string]interface{}{
		"label":              label,
		"name":               getField(&e, label, "Name"),
		"registrationNumber": getField(&e, label, "RegistrationNumber"),
		"email":              getField(&e, label, "Email"),
		"phoneNumber":        getField(&e, label, "PhoneNumber"),
		"gender":             getField(&e, label, "Gender"),
	})
	if err != nil {
		c <- err
	}
	result.Next()
	log.Println(result.Record().GetByIndex(0))

	if err = result.Err(); err != nil {
		c <- err
	}
	log.Printf("Created %s node", label)
	c <- nil
}
