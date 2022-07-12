package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HelloHandler struct {
	Message string `json:"message"`
}

func New(msg string) *HelloHandler {
	return &HelloHandler{
		Message: msg,
	}
}

func (app *HelloHandler) HelloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	/* data := struct {
		Message string `json:"message"`
	}{
		Message: app.Message,
	} */
	json.NewEncoder(w).Encode(app)

}

func (app *HelloHandler) CurrentTimeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format(time.Kitchen)
	w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}
