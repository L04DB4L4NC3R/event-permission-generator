package main

import (
	"log"
	"net/http"
	"text/template"

	"./controller"
	"./model"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func main() {
	// connect to database
	session, driver, result, err := connectToDB()
	if err != nil {
		log.Fatalln("Error connecting to Database")
		log.Fatalln(err)
	}
	log.Println("Connected to Neo4j")
	// Close driver and session after func ends
	defer driver.Close()
	defer session.Close()

	// pass the session to the model layer
	model.SetDB(session, result)

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

func connectToDB() (neo4j.Session, neo4j.Driver, *neo4j.Result, error) {

	// define driver, session and result vars
	var (
		driver  neo4j.Driver
		session neo4j.Session
		result  *neo4j.Result
		err     error
	)

	// initialize driver to connect to localhost with ID and password
	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("angad", "angad", "")); err != nil {
		return nil, nil, nil, err
	}

	// Open a new session with write access
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, nil, nil, err
	}
	return session, driver, result, nil
}
