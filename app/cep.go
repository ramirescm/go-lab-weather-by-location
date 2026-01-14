package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var ErrZipNotFound = errors.New("zip not found")

type viaCepResponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro"`
}

func GetCityByCEP(cep string) (string, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("viacep error")
	}

	var data viaCepResponse
	json.NewDecoder(resp.Body).Decode(&data)

	if data.Erro {
		return "", ErrZipNotFound
	}

	return data.Localidade, nil
}
