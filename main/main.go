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
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./doc"))))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func populateTempaltes() *template.Template {
	res := template.New("data.html")
	template.Must(res.ParseGlob("./static/*.html"))
	return res
}
