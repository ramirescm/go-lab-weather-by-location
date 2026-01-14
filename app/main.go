package main

import (
	"log"
	"net/http"
)

func main() {

	port := "8080"

	http.HandleFunc("/weather", WeatherHandler)

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
