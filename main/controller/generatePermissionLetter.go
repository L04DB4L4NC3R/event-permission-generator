package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	events "../../lib"
)

type letter struct {
	temp *template.Template
}

func Handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

/**
 * @api {post} /permissionLetter generate permission letter
 * @apiName permissionLetter
 * @apiGroup all
 *
 * @apiParam {String} clubName Name of your club
 * @apiParam {String} name Name of your event
 * @apiParam {String} toDate ending date
 * @apiParam {String} fromDate start date
 * @apiParam {String} toTime start time
 * @apiParam {String} fromTime ending time
 * @apiParam {String} budget budget
 * @apiParam {String} description event description
 * @apiParam {String} [category] category of the event
 * @apiParam {String} venue event venue
 * @apiParam {String} [attendance] Name of your club
 * @apiParam {String} expectedParticipants Expected Participants in the event
 * @apiParam {Object} facultyCoordinator faculty coordinator details (Participant Object)
 * @apiParam {Object} studentCoordinator student coordinator details (Participant Object)
 * @apiParam {Object} guest guest details (Guest object)
 * @apiParam {String} PROrequest PRO department requests
 * @apiParam {String} campusEngineerRequest Campus engineer requests
 * @apiParam {String} duration duration of event
 * @apiParam {Object} mainSponsor sponsor details (Participant Object)
 *
 * @apiParamExample {json} request-example
 *
 * {
   "clubName":"GDG",
   "name":"DEVFEST",
   "toDate":"10TH OCTOBER",
   "fromDate":"8TH OCTOBER",
   "toTime":"10 PM",
   "fromTime":"11 AM",
   "budget":"200000",
   "description":"TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING",
   "category":"TECHNICAL",
   "venue":"ANNA AUDI",
   "attendance":"4000",
   "expectedParticipants":"4000",
   "facultyCoordinator":{
      "name":"NOORU MAA",
      "registrationNumber":"17BBE1010",
      "email":"SDADAS@A.COM",
      "phoneNumber":"919191991911",
      "gender":"M",
      "eventsAttended":"ALL"
   },
   "studentCoordinator":{
      "name":"NOORU BAAP",
      "registrationNumber":"17BBE1010",
      "email":"SDADAS@A.COM",
      "phoneNumber":"919191991911",
      "gender":"M",
      "eventsAttended":"ALL"
   },
   "guest":{
      "name":"ALLAHH DAAS",
      "email":"ASDSAD#ASD.COM",
      "phoneNumber":"11111111111",
      "gender":"F",
      "stake":"SOME MONAYYYY",
      "locationOfStay":"TERA GHAR"
   },
   "PROrequest":"SAJDOOSIJANDFSAKFDSAFD",
   "campusEngineerRequest":"SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
   "duration":"16 hours",
   "mainSponsor":{
      "name":"ALLAHH DAAS",
      "email":"ASDSAD#ASD.COM",
      "phoneNumber":"11111111111",
      "gender":"F",
      "stake":"SOME MONAYYYY",
      "locationOfStay":"TERA GHAR"
   }
}
 *
 * @apiSuccess {html} html Get html.
*/
func (l letter) handleLetter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data events.Event
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
