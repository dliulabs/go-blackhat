package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	body := struct {
		foo string `json:"foo"`
	}{
		foo: "bar",
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	request, err := http.NewRequest(
		"PUT",
		"https://postman-echo.com/put",
		b)
	if err != nil {
		panic(err)
	}
	var client http.Client
	resp, err := client.Do(request)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

}
