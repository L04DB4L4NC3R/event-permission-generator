package model

import "log"

func CreateEvent(e Event) {
	result, err := session.Run("CREATE (n:EVENT) {name:$name, clubName $clubName} RETURN n", map[string]interface{}{
		"name":     e.Name,
		"clubName": e.ClubName,
	})
	handleErr(err)

	for result.Next() {
		log.Println(result.Record().GetByIndex(0).(string))
	}

}
