package main

import (
	"log"
	"net/http"

	"./controller"
)

func main() {
	controller.Startup()
	log.Println("Starting to listen..")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
