package main

import (
	"encoding/json"
	"fmt"
	"log"
	"openweather/openweather"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: main <city> <state>")
	}
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalln("You must define OPENWEATHER_API_KEY env variable.")
	}
	s := openweather.New(apiKey)

	citySearch, err := s.WeatherSearch(os.Args[1], os.Args[2])
	if err != nil {
		log.Panicln(err)
	}
	j, _ := json.Marshal(citySearch)
	fmt.Printf("%v\n", string(j))
}
