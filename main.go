package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting to listen..")
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./doc"))))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// package main

// import (
// 	"log"
// 	"reflect"

// 	"./main/model"
// )

// func getField(v *model.Event, field string) reflect.Value {
// 	r := reflect.ValueOf(v)
// 	f := reflect.Indirect(r).FieldByName(field)
// 	return f
// }

// func main() {

// 	var e = model.Event{FacultyCoordinator: model.Participant{
// 		Name: "Angad Sharma",
// 	}}
// 	log.Println(getField(&e, "FacultyCoordinator").FieldByName("Name"))

// }
