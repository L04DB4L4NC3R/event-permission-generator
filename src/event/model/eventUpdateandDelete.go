package model

import (
	"log"

	events "../../../lib"
)

func DeleteEvent(q Query, c chan error) {
	result, err := events.Session.Run(`
		MATCH(n:EVENT)-[r]->(a)
		WHERE n.`+q.Key+`=$val
		DETACH DELETE n, a
	`, map[string]interface{}{
		"val": q.Value,
	})
	if err != nil {
		c <- err
	}

	result.Next()
	log.Println(result.Record())

	if err = result.Err(); err != nil {
		c <- err
	}
	log.Println("Event deleted")
	c <- nil
}
