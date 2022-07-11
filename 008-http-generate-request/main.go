package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	body := struct {
		Foo string `json:"foo"` // be sure the fileds are in upper case
	}{
		Foo: "bar",
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)
	request, err := http.NewRequest(
		"PUT",
		"http://postman-echo.com/put",
		b)
	if err != nil {
		panic(err)
	}
	var client http.Client
	resp, err := client.Do(request)
	if err != nil {
		log.Panicln(err)
	}
	// resp.Body is an io.ReadCloser
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(respData))
	resp.Body.Close()
	fmt.Println(resp.Status)

}
