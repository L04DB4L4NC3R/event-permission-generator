package controller

import (
	"encoding/json"
	"log"
	"net/http"

	
	events "github.com/angadsharma1016/omega_dbconfig"
)

/**
 * @api {post} /event/create create a new event
 * @apiName create a new event
 * @apiGroup admin
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
* @apiParamExample {json} response-example
{
	"message":true
}
 * @apiSuccess {Status} Boolean true or false.
*/
func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var data events.Event
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	ce := make(chan error)
	go model.CreateEvent(data, ce)

	if err = <-ce; err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{false, "some error occurreed"})
		return
	}

	json.NewEncoder(w).Encode(struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}{true, "new node created successfully"})
}

/**
 * @api {get} /event/read read an event
 * @apiName read an event
 * @apiGroup admin
 *
 * @apiParam {String} key key to query the event by
 * @apiParam {String} value value of the key

@apiParamExample {json} request-example
{
	"key":"name",
	"value":"DEVFEST"
}

 * @apiParamExample {json} response-example
 *
 * {
    "clubName": "GDG",
    "name": "DEVFEST",
    "toDate": "10TH OCTOBER",
    "fromDate": "8TH OCTOBER",
    "toTime": "10 PM",
    "fromTime": "11 AM",
    "budget": "200000",
    "description": "TECHNICAL EVENT AT GDG VIT. ITS GONNA BE AMAZING",
    "category": "TECHNICAL",
    "venue": "ANNA AUDI",
    "attendance": "4000",
    "expectedParticipants": "4000",
    "facultyCoordinator": {
        "name": "NOORU MAA",
        "registrationNumber": "17BBE1010",
        "email": "SDADAS@A.COM",
        "phoneNumber": "919191991911",
        "gender": "M",
        "eventsAttended": ""
    },
    "studentCoordinator": {
        "name": "NOORU BAAP",
        "registrationNumber": "17BBE1010",
        "email": "SDADAS@A.COM",
        "phoneNumber": "919191991911",
        "gender": "M",
        "eventsAttended": ""
    },
    "guest": {
        "name": "ALLAHH DAAS",
        "email": "ASDSAD#ASD.COM",
        "phoneNumber": "11111111111",
        "gender": "F",
        "stake": "SOME MONAYYYY",
        "locationOfStay": "TERA GHAR"
    },
    "PROrequest": "SAJDOOSIJANDFSAKFDSAFD",
    "campusEngineerRequest": "SDFHBSADUB, ASNFD , AS KDFSAM FDSA, AS, SD",
    "duration": "16 hours",
    "mainSponsor": {
        "name": "",
        "registrationNumber": "",
        "email": "",
        "phoneNumber": "",
        "gender": "",
        "eventsAttended": ""
    }
}
 *
*/
func readEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var query model.Query
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		log.Println(err)
	}

	c := make(chan model.EventReturn)
	go model.ShowEventData(query, c)

	var ret model.EventReturn = <-c
	if ret.Err != nil {
		log.Println(ret.Err)
		json.NewEncoder(w).Encode(struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{false, "some error occurreed"})
		return
	}
	json.NewEncoder(w).Encode(ret.Event)

}

/**
 * @api {delete} /event/delete delete an event
 * @apiName delete an event
 * @apiGroup admin
 *
 * @apiParam {String} key key to query the event by
 * @apiParam {String} value value of the key

@apiParamExample {json} request-example
{
	"key":"name",
	"value":"DEVFEST"
}

 * @apiParamExample {json} response-example
 *
 * {"status":true,"message":"node DELETED successfully"}

 *
*/
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var data model.Query
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	c := make(chan error)
	go model.DeleteEvent(data, c)

	if err = <-c; err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{false, "some error occurreed"})
		return
	}

	json.NewEncoder(w).Encode(struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}{true, "node DELETED successfully"})
}

/**
 * @api {put} /event/update update an event
 * @apiName update an event
 * @apiGroup admin
 *
 * @apiParam {String} key key to query the event by
 * @apiParam {String} value value of the key
 * @apiParam {String} changeKey key of the value which needs to be altered
 * @apiParam {String} changeValue the new value

@apiParamExample {json} request-example
{
	"key":"name",
	"value":"DEVRELCONF",
	"changeKey":"some key",
	"changeValue":"some value"
}

 * @apiParamExample {json} response-example
 *
 * {"status":true,"message":"Updated name to DEVRELCONF"}

 *
*/
func updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data model.Query
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
	}

	c := make(chan error)
	go model.UpdateEvent(data, c)

	if err = <-c; err != nil {

		log.Println(err)
		json.NewEncoder(w).Encode(struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{false, "some error occurreed"})
		return
	}
	log.Println("hagga")

	json.NewEncoder(w).Encode(struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}{true, "node updated successfully"})
}

func eventCRUDHandler() {
	http.HandleFunc("/api/v1/event/create", createEvent)
	http.HandleFunc("/api/v1/event/read", readEvent)
	http.HandleFunc("/api/v1/event/delete", deleteEvent)
	http.HandleFunc("/api/v1/event/update", updateEvent)
}
