package main

import (
	apifarm "apifarm/src"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var storage = apifarm.NewInMemory()
var service = apifarm.NewVideoGameService(storage)
var controller = apifarm.NewController(service)

// api testing entities
var dataLoader = apifarm.NewJSONFileLoader(storage)
var apiTestingController = apifarm.NewAPITestingController(dataLoader)

func main() {
	env := os.Getenv("API_ENV")

	if env != "PROD" {
		env = "DEV"
	}

	p := fmt.Sprintf("./config.%s.json", strings.ToLower(env))

	c, err := apifarm.GetConfiguration(p)

	if err != nil {
		panic(err)
	}

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

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	log.Fatal(http.ListenAndServe(addr, r))
}
