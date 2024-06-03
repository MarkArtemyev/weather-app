package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string) (WeatherData, error) {
	var data WeatherData
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("error: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	if os.Getenv("API_KEY") == "" {
		fmt.Println("Error: API_KEY is not set in the environment")
		os.Exit(1)
	}

	var city string
	fmt.Print("Enter city: ")
	fmt.Scanln(&city)

	data, err := getWeather(city)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Weather in %s:\n", data.Name)
	fmt.Printf("Temperature: %.2f°C\n", data.Main.Temp)
	fmt.Printf("Description: %s\n", data.Weather[0].Description)
}
