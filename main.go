package main

import (
	"goordinater/controller"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	cont := controller.CoordinateController()
	mux.Handle(cont.Path, http.HandlerFunc(cont.Handle))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
