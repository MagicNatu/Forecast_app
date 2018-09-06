package main

// Defining structs to store recieved JSON data in.

import (
)

type ForecastWeatherData struct {
	Unit    string
	Lang    string
	Key     string
	baseURL string
}

type forecast struct {
	Days int `json:"cnt"`
	Dt int `json:"dt"`
	List []ForecastWeatherList `json:"list"`
}

type ForecastWeatherList struct {
	Dt      int       `json:"dt"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Date	string	  `json:"dt_txt"`
}

type Main struct {
	Temp_min float64  `json:"temp_min"`
	Temp_max float64  `json:"temp_max"`
	Avg_Temp float64 `json:"temp"`
}

type Weather struct{
	Sky string `json:"description"`
}

func NewWeatherData(unit string, lang string) ForecastWeatherData {
	data := ForecastWeatherData{}
	data.Key = apiKey
	data.Lang = lang
	data.Unit = unit
	data.baseURL = "http://api.openweathermap.org/data/2.5/forecast/?q="
	return data
}