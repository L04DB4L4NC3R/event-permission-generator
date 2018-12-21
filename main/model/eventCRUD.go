package model

import "log"

func CreateEvent(e Event) error {

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

	// for result.Next() {
	// 	log.Println(result.Record().GetByIndex(1))
	// }
	result.Next()
	log.Println(result.Record().GetByIndex(0).(string))

	if err = result.Err(); err != nil {
		return err
	}

	return nil
}
