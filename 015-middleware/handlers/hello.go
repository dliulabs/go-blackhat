package handlers

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (app *HelloHandler) HelloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}
