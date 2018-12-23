package model

import "log"

type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EventReturn struct {
	Event Event
	Err   error
}

func ShowEventData(q Query, c chan EventReturn) {
	result, err := session.Run(`
		MATCH(n:EVENT) 
		WHERE n.`+q.Key+`=$val
		RETURN n.clubName, n.name, n.toDate, n.fromDate, n.toTime, n.fromTime, n.budget, n.description, n.category,
		n.venue, n.attendance, n.expectedParticipants, n.PROrequest, n.campusEngineerRequest, n.duration
	`, map[string]interface{}{
		"val": q.Value,
	})

	if err != nil {
		c <- EventReturn{Event{}, err}
	}
	var ev Event
	for result.Next() {
		ev = Event{
			ClubName:              result.Record().GetByIndex(0).(string),
			Name:                  result.Record().GetByIndex(1).(string),
			ToDate:                result.Record().GetByIndex(2).(string),
			FromDate:              result.Record().GetByIndex(3).(string),
			ToTime:                result.Record().GetByIndex(4).(string),
			FromTime:              result.Record().GetByIndex(5).(string),
			Budget:                result.Record().GetByIndex(6).(string),
			Description:           result.Record().GetByIndex(7).(string),
			Category:              result.Record().GetByIndex(8).(string),
			Venue:                 result.Record().GetByIndex(9).(string),
			Attendance:            result.Record().GetByIndex(10).(string),
			ExpectedParticipants:  result.Record().GetByIndex(11).(string),
			PROrequest:            result.Record().GetByIndex(12).(string),
			CampusEngineerRequest: result.Record().GetByIndex(13).(string),
			Duration:              result.Record().GetByIndex(14).(string),
		}
	}

	if err = result.Err(); err != nil {
		c <- EventReturn{ev, err}
	}
	log.Println(ev)
	c <- EventReturn{ev, nil}
}
