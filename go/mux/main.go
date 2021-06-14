package main

import (
	apifarm "apifarm/src"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var storage = apifarm.NewInMemory()
var service = apifarm.NewVideoGameService(storage)
var controller = apifarm.NewController(service)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		controller.HandlePing(apifarm.NewHTTPResponse(&w))
	})

	r.HandleFunc("/video_games", func(w http.ResponseWriter, r *http.Request) {
		controller.HandleGetAll(apifarm.NewHTTPResponse(&w))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
