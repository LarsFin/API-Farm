package main

import (
	apifarm "apifarm/src"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))

	if err != nil {
		panic(err)
	}
}

func squareHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ns := vars["number"]

	i := strToInt(ns)

	i = apifarm.Square(i)

	o := strconv.Itoa(i)

	_, err := w.Write([]byte(o))

	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/square/{number:[0-9]+}", squareHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}
