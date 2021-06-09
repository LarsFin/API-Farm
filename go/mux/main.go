package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
