package main

import (
	apifarm "apifarm/src"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var controller = apifarm.Controller{}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		controller.HandlePing(apifarm.NewHttpResponse(&w))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}
