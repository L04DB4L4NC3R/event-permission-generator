package model

import "log"

type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Label string `json:"label"`
}

func ShowData(q Query) error {
	result, err := session.Run(`
		MATCH(n:`+q.Label+`) 
		WHERE n.`+q.Key+`=$val
		RETURN n
	`, map[string]interface{}{
		"val": q.Value,
	})

	if err != nil {
		return err
	}

	for result.Next() {
		log.Printf("%+v\n", result.Record().GetByIndex(0))
	}

	if err = result.Err(); err != nil {
		return err
	}
	return nil
}
