package openweather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	apiKey string
}

const BaseURL = "https://api.openweathermap.org"

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

type Weather struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Result struct {
	Result Weather `json:"main"`
}

func (s *Client) WeatherSearch(city string, state string) (*Result, error) {
	res, err := http.Get(
		// fmt.Sprintf("%s/data/2.5/weather?q=%s&APPID=%s", BaseURL, q, s.apiKey),
		fmt.Sprintf("%s/data/2.5/weather?q=%s,%s&APPID=%s", BaseURL, city, state, s.apiKey),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	log.Printf(res.Status)

	var ret Result
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
