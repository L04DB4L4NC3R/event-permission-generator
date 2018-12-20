package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var data model.Event
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}
	model.CreateEvent(data)

	json.NewEncoder(w).Encode(struct {
		message string
	}{"Done"})
}

func eventCRUDHandler() {
	http.HandleFunc("/api/v1/event/create", createEvent)
}
