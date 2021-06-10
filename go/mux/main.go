package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))

	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
