package main

import (
	"log"
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"strconv"
)

// Defining key environment variable in order to access weather data from openweathermap.org

var setenv = os.Setenv("OWM_API_KEY", "ae109bfa9bd34e64691b5b467fe631eb")
var apiKey = os.Getenv("OWM_API_KEY")

// Retrieves forecast data from openweathermap.org, takes location and amount of days as arguments. Structs defined in forecast.go

func GetForecast(location string, days int) forecast {
	var f ForecastWeatherData
	f = NewWeatherData("metric", "EN")
	actual_days := (days*8)+3
	days_string := strconv.Itoa(actual_days)
	response, err := http.Get(fmt.Sprintf("%s%s&appid=%s&units=%s&lang=%s&cnt=%s",f.baseURL,location,f.Key, f.Unit,f.Lang, days_string))
	 if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	
	forecastData := forecast{}
	err = json.NewDecoder(response.Body).Decode(&forecastData)
	if err != nil {
        panic(err)
    }
	
	return forecastData
}

// retrieves current weather data from openweathermap.org, returns a current struct. Structs defined in currentWeather.go.
 
func GetCurrent(location string) Current {
	var c CurrentWeatherData
	c = NewCurrentWeatherData("metric", "EN")
	response, err := http.Get(fmt.Sprintf("%s%s&appid=%s&units=%s&lang=%s&cnt=%s",c.baseURL,location,c.Key, c.Unit,c.Lang))
	 if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	
	currentData := Current{}
	err = json.NewDecoder(response.Body).Decode(&currentData)
	if err != nil {
        panic(err)
    }
	return currentData
}

// Function to flush variables in order to avoid bugs in logline

func flush(f string) {
	f = ""
	fmt.Scanf("%s", &f)
}

// Wraps functionality by printing the requested forecast and current weather data

func printOutput() {
	var location string
	var max, min, amt_days, frequency string

	for true {
		fmt.Print("Enter a city, or exit by typing exit: ")
		fmt.Scanf("%s", &location)
		if location == "exit" {
			os.Exit(0)
	}
		flush(location)
		fmt.Print("For how many days do you want to monitor the temp? (max 5) ")
		fmt.Scanf("%s", &amt_days)
		
		amt_days_int, err := strconv.ParseInt(amt_days, 10, 32) 
			if err != nil {
			log.Fatalln(err)
	}
		if amt_days_int > 5 {
			fmt.Println("Max 5!!")
			continue
		}
		flush(amt_days)
		fmt.Print("Enter max temp: ")
		fmt.Scanf("%s", &max)
		flush(max)
		fmt.Print("Enter min temp: ")
		fmt.Scanf("%s", &min)
		flush(min)
		fmt.Print("What is the frequency of check? Valid inputs: 3,6,9,12,15,18,21,24 (In hours)")
		fmt.Scanf("%s", &frequency)
		frequency_int, err := strconv.ParseInt(frequency, 10, 32) 
			if err != nil {
			log.Fatalln(err)
		}
			flush(frequency)
			
			if (frequency_int % 3) == 0 && frequency_int <= 24 {
			} else {
				fmt.Println("Invalid input")
				continue
			}		
		 	if location != "" {
			var c Current = GetCurrent(location)
			fmt.Printf("Current temp for %s is %f degrees celcius\n\n",location, c.Main.CurrentTemp)
			
			days, err := strconv.Atoi(amt_days)
			if err != nil {
			log.Fatalln(err)	
	}
			
			var f forecast = GetForecast(location, days)
			
			min_float64, err := strconv.ParseFloat(min, 64) 
			if err != nil {
			log.Fatalln(err)
		
	}
			max_float64, err := strconv.ParseFloat(max, 64) 
			if err != nil {
			log.Fatalln(err)
		
	}
			
			fmt.Println("The forecast data for the next", f.Days/8, "day(s) accoring to max temp and min temp is as follows:\n")
			var freq int
			freq = int(frequency_int)
			freq = freq/3
			
			for i := 1; i<f.Days; i = i+freq {
				if f.List[i].Main.Temp_min < min_float64 {
				fmt.Printf("Temperature for %v is minimum: %f degrees and is less than specified min temp %s\n", f.List[i].Date, f.List[i].Main.Temp_min, min)
				}
				if f.List[i].Main.Temp_max > max_float64 {
				fmt.Printf("Temperature for %v is maximum: %f degrees and is more than specified max temp %s\n", f.List[i].Date, f.List[i].Main.Temp_max, max)
				}
				if (f.List[i].Main.Temp_min >= min_float64) && (f.List[i].Main.Temp_max <= max_float64) {
					fmt.Printf("Min & max temperature for %v is within specified limits; average temp is %f: \n", f.List[i].Date, f.List[i].Main.Avg_Temp)
				}
			}
			
			flush(location)
			
		}
	}
}

func main() {
	printOutput()
}