package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temp       float64 `json:"temp"`
	Feels_like float64 `json:"feels_like"`
	TempMax    float64 `json:"temp_max"`
	TempMin    float64 `json:"temp_min"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Gust  float64 `json:"gust"`
}

type Sys struct {
	Country string `json:"country"`
}

type WeatherData struct {
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Sys     Sys       `json:"sys"`
	Name    string    `json:"name"`
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Retrieve API_KEY from environment variables
	apiKey := os.Getenv("KEY")
	if apiKey == "" {
		fmt.Println("API_KEY not found in environment variables")
		return
	}
	//URL for api
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=44.0128&lon=20.9114&appid=%s&units=metric", apiKey)

	//Get request to api
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather data: %v", err)
		return
	}

	defer response.Body.Close()

	//Api response returns data in bytes, which is why we are using io.ReadAll to read the response body into a byte slice.
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)
		return
	}
	//When you print body directly using fmt.Println(body), it prints the byte slice body as a sequence of numbers representing the byte values.
	//When you use fmt.Println(string(body)), you are converting the byte slice body to a string using the string() function
	// fmt.Println(string(body))

	var weatherData WeatherData

	//Unmarshal JSON data into the weatherData variable
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON data: %v", err)
		return
	}

	// fmt.Println(weatherData)
	now := time.Now()
	fmt.Printf("%s,%s %s %s %s \nDescription: %s\nTemp: %.0fC \nWind speed: %0.0f m/s\n", weatherData.Name, weatherData.Sys.Country, now.Format("2006-01-02"), now.Format("15:04:05"), weatherData.Weather[0].Main, weatherData.Weather[0].Description, weatherData.Main.Temp, weatherData.Wind.Speed)

}
