package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var zipRegex = regexp.MustCompile(`^\d{8}$`)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")
	if apiKey == "" {
		http.Error(w, "API key missing", http.StatusUnauthorized)
		return
	}

	cep := r.URL.Query().Get("cep")
	if !zipRegex.MatchString(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := GetCityByCEP(cep)
	if err == ErrZipNotFound {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	fmt.Println(err)
	if err != nil {
		http.Error(w, "1 internal error", http.StatusInternalServerError)
		return
	}

	tempC, err := GetTemperature(city, apiKey)
	if err != nil {
		http.Error(w, "2 internal error "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempC*1.8 + 32,
		"temp_K": tempC + 273,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
