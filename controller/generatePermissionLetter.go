package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../helpers"
	"../model"
)

func Handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func handleLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data model.Event
		err := json.NewDecoder(r.Body).Decode(&data)

		Handle(err)
		w.Write([]byte(helpers.GenerateLetter(data)))
	}
}

func permissionLetterHandler() {
	http.HandleFunc("/api/v1/permissionLetter", handleLetter)
}
