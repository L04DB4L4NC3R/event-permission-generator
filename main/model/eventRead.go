package model

import (
	"log"
)

type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ShowEventData(q Query) error {
	result, err := session.Run(`
		MATCH(n:EVENT) 
		WHERE n.`+q.Key+`=$val
		RETURN n.clubName, n.name, n.toDate, n.fromDate, n.toTime, n.fromTime, n.budget, n.description, n.category,
		n.venue, n.attendance, n.expectedParticipants, n.PROrequest, n.campusEngineerRequest, n.duration
	`, map[string]interface{}{
		"val": q.Value,
	})

	if err != nil {
		return err
	}

	var ev

	for result.Next() {
		ev := Event {
			ClubName              :      result.Record().GetByIndex()
			Name                  :      result.Record().GetByIndex()
			ToDate                :      result.Record().GetByIndex()
			FromDate              :      result.Record().GetByIndex()
			ToTime                :      result.Record().GetByIndex()
			FromTime              :      result.Record().GetByIndex()
			Budget                :      result.Record().GetByIndex()
			Description           :      result.Record().GetByIndex()
			Category              :      result.Record().GetByIndex()
			Venue                 :      result.Record().GetByIndex()
			Attendance            :      result.Record().GetByIndex()
			ExpectedParticipants  :      result.Record().GetByIndex()
			PROrequest            :      result.Record().GetByIndex()
			CampusEngineerRequest :      result.Record().GetByIndex()
			Duration              :      result.Record().GetByIndex()
		}
	}

	if err = result.Err(); err != nil {
		return err
	}

	return nil
}
