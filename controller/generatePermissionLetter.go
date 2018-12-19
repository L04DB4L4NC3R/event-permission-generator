package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"../model"
)

type letter struct {
	temp *template.Template
}

func Handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (l letter) handleLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data model.Event
		err := json.NewDecoder(r.Body).Decode(&data)

		Handle(err)
		t := l.temp.Lookup("data.html")
		if t != nil {
			err = t.Execute(w, data)
			Handle(err)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	}
}

func (l letter) permissionLetterHandler() {
	http.HandleFunc("/api/v1/permissionLetter", l.handleLetter)
}
