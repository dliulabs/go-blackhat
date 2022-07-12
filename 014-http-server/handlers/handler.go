package handlers

import (
	"encoding/json"
	"net/http"
)

//Logger is a middleware handler that does request logging
type ApiHandler struct{}

//NewLogger constructs a new Logger middleware handler
func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		data := struct {
			Message string `json:"message"`
		}{
			Message: "Hello, World",
		}
		json.NewEncoder(w).Encode(data)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}
