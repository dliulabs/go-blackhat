package main

import (
	"hellomw/handlers"
	"net/http"
)

func main() {
	app := handlers.NewHelloHandler()
	addr := "0.0.0.0:8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", app.HelloHandlerFunc)
	/* bare Server
	// http.ListenAndServe(addr, mux)
	*/

	/* middleware */
	logger := handlers.NewLoggerHandler(mux)
	http.ListenAndServe(addr, logger)
}
