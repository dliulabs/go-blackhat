package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "hi foo\n")
	}).Methods("GET")
	router.HandleFunc("/bar/{bar}", func(w http.ResponseWriter, req *http.Request) {
		bar := mux.Vars(req)["bar"] // mux.Vars() returns map[string]string
		fmt.Fprintf(w, "hi %s\n", bar)
	}).Methods("GET")
	// using regex
	router.HandleFunc("/users/{user:[a-z]+}", func(w http.ResponseWriter, req *http.Request) {
		user := mux.Vars(req)["user"]
		fmt.Fprintf(w, "hi %s\n", user)
	}).Methods("GET")

	ADDR := ":8080"
	log.Printf("Listening at %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, router))
}
