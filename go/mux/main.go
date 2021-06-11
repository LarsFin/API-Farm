package main

import (
	apifarm "apifarm/src"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = apifarm.Controller{}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		controller.HandlePing(apifarm.NewHTTPResponse(&w))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
