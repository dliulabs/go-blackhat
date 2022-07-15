package main

import (
	"hellomux/handlers"
	"log"
	"net/http"
)

var helloApp *handlers.HelloHandler

func init() {
	helloApp = handlers.New("Hello, World")
}

func main() {
	addr := "0.0.0.0:8080"

	mux := http.NewServeMux() // an http.Handler
	mux.HandleFunc("/v1/hello", helloApp.HelloHandlerFunc)
	mux.HandleFunc("/v1/time", helloApp.CurrentTimeHandlerFunc)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
