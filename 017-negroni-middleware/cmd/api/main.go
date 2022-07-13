package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"negronimw/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "hi foo\n")
	}).Methods("GET")

	n := negroni.Classic()
	{
	// Use adds a Handler onto the middleware stack. Handlers are invoked in the order they are added to a Negroni.
	// this means handlers.ServeHttp() will be used
		n.Use(&handlers.BasicAuth{
			Username: "admin",
			Password: "password",
		})
	}

	n.UseHandler(r)

	ADDR := ":8080"
	log.Printf("Listening at %s", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, n)) 
}
