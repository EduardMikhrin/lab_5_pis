package main

const (
	MoodleLogin = "is-33-fiot-126"
	Surname     = "Mikhrin"
	Name        = "Eduard"
	Course      = 2
	Group       = "IS-33"
)

type WeatherResponse struct {
	CurrentCondition []struct {
		TempC         string `json:"temp_C"`
		FeelsLikeC    string `json:"FeelsLikeC"`
		WindspeedKmph string `json:"windspeedKmph"`
		Humidity      string `json:"humidity"`
	} `json:"current_condition"`
}

type ViewData struct {
	Date          string
	TempC         string
	FeelsLikeC    string
	WindspeedKmph string
	Humidity      string
}
