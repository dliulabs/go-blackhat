package handlers

import (
	"log"
	"net/http"
)

type Trivial struct {
}

func (t *Trivial) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Executing trivial middleware")
	next(w, r)
}
