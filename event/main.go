package main

import (
	"log"
	"net/http"
	"text/template"

	"./controller"
	events "github.com/angadsharma1016/omega_dbconfig"
)

func main() {
	// connect to database
	session, driver, err := events.ConnectToDB()
	if err != nil {
		log.Fatalln("Error connecting to Database")
		log.Fatalln(err)
	}
	log.Println("Connected to Neo4j")
	// Close driver and session after func ends
	defer driver.Close()
	defer session.Close()

	// pass the session to the model layer
	events.SetDB(session)

	// populate templates
	controller.Startup(populateTempaltes())

	// listen on specified port
	log.Println("Starting to listen..")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func populateTempaltes() *template.Template {
	res := template.New("data.html")
	template.Must(res.ParseGlob("./static/*.html"))
	return res
}
