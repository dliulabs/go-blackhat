package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ReqData struct {
	Foo string `json:"data"`
}

type RespData struct {
	Data ReqData `json:"data"`
}

func main() {
	body := ReqData{
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
	defer resp.Body.Close()
	// resp.Body is an io.ReadCloser
	//respData, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// fmt.Println(string(respData))
	fmt.Println(resp.Status)

	var data RespData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%v\n", data)
}
