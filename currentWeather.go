package main

// Defining structs to store recieved JSON data in.

import (
)
type CurrentWeatherData struct {
	Unit    string
	Lang    string
	Key     string
	baseURL string
}
type Current struct {
	Main CurrentWeather `json:"main"`
}
type CurrentWeather struct {
	CurrentTemp float64 `json:"temp"`
}
func NewCurrentWeatherData(unit string, lang string) CurrentWeatherData {
	data := CurrentWeatherData{}
	data.Key = apiKey
	data.Lang = lang
	data.Unit = unit
	data.baseURL = "http://api.openweathermap.org/data/2.5/weather/?q="
	return data
}
