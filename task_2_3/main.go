package main

import (
	"log"
	"net/http"
)

func main() {
	// Маршрути
	http.HandleFunc("/", handlerWeather)
	http.HandleFunc("/"+MoodleLogin, handlerProfile)

	log.Println("Сервер запущено на http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
