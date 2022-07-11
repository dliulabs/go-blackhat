package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	form := url.Values{}
	form.Add("foo", "bar")
	resp, err := http.Post(
		"https://postman-echo.com/post",
		//"https://www.google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	// fmt.Println(resp.Status)
	// Read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()
}
