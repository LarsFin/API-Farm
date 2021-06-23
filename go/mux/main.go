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

// api testing entities
var dataLoader = apifarm.NewJSONFileLoader(storage)
var apiTestingController = apifarm.NewAPITestingController(dataLoader)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		controller.HandlePing(apifarm.NewHTTPResponse(&w))
	})

	r.HandleFunc("/video_games/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.HandleGet(apifarm.NewHTTPRequest(r), apifarm.NewHTTPResponse(&w))
	}).Methods(http.MethodGet)

	r.HandleFunc("/video_games", func(w http.ResponseWriter, r *http.Request) {
		controller.HandleGetAll(apifarm.NewHTTPResponse(&w))
	}).Methods(http.MethodGet)

	r.HandleFunc("/video_games", func(w http.ResponseWriter, r *http.Request) {
		controller.HandlePost(apifarm.NewHTTPRequest(r), apifarm.NewHTTPResponse(&w))
	}).Methods(http.MethodPost)

	r.HandleFunc("/video_games/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.HandlePut(apifarm.NewHTTPRequest(r), apifarm.NewHTTPResponse(&w))
	}).Methods(http.MethodPut)

	r.HandleFunc("/api_tests/setup", func(w http.ResponseWriter, r *http.Request) {
		apiTestingController.HandleTestSetup(apifarm.NewHTTPResponse(&w))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
