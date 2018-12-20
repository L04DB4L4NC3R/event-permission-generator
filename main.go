package main

import (
	"log"
	"net/http"
	"text/template"

	"./controller"
	_ "github.com/swaggo/http-swagger"
)

// @title Project Omega documentation
// @version 0.1
// @description The last event manager you will ever need
// @host localhost:3000
// @BasePath /api/v1
func main() {
	controller.Startup(populateTempaltes())
	log.Println("Starting to listen..")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func populateTempaltes() *template.Template {
	res := template.New("data.html")
	template.Must(res.ParseGlob("./static/*.html"))
	return res
}
