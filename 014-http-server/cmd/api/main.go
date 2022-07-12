package main

import (
	"hellosvc/handlers"
	"log"
	"net/http"
)

var apiApp *handlers.ApiHandler

func init() {
	apiApp = handlers.NewApiHandler()
}

func main() {
	addr := "0.0.0.0:9090"
	log.Printf("server is listening at %s", addr)
	//use wrappedMux instead of mux as root handler
	log.Fatal(http.ListenAndServe(addr, apiApp))
}
