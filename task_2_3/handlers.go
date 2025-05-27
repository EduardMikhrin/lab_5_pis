package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func handlerWeather(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://wttr.in/Kyiv?format=j1")
	if err != nil {
		http.Error(w, "Не вдалося отримати дані: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Помилка читання відповіді: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var data WeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Помилка розбору JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	wc := data.CurrentCondition[0]
	vd := ViewData{
		Date:          time.Now().Format("2006-01-02 15:04:05"),
		TempC:         wc.TempC,
		FeelsLikeC:    wc.FeelsLikeC,
		WindspeedKmph: wc.WindspeedKmph,
		Humidity:      wc.Humidity,
	}

	if err := tpl.Execute(w, vd); err != nil {
		log.Println("Помилка рендеру шаблону:", err)
	}
}

func handlerProfile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"+MoodleLogin {
		http.NotFound(w, r)
		return
	}
	info := struct {
		Surname string `json:"surname"`
		Name    string `json:"name"`
		Course  int    `json:"course"`
		Group   string `json:"group"`
	}{
		Surname,
		Name,
		Course,
		Group,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(info)
}
