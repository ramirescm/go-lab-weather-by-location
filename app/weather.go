package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func GetTemperature(city, key string) (float64, error) {
	cityEncoded := url.QueryEscape(city)
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + cityEncoded + "&appid=" + key + "&units=metric"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("openweathermap error")
	}

	var data WeatherResponse
	json.NewDecoder(resp.Body).Decode(&data)
	return data.Main.Temp, nil
}
