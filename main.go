package main

import (
	"log"
	"net/http"
	"text/template"

	"./controller"
)

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
