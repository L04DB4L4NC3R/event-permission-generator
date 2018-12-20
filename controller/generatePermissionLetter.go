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

// handleLetter godoc
// @Summary Generate a permission letter for the event
// @Description Generate a permission letter for the event
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Param name path string true "Name of the event"
// @Success 200 {object} model.Event
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /accounts/{id} [get]
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
