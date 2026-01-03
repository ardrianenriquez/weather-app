package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// https://api.openweathermap.org/data/2.5/weather?q={city name}&appid={API key}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL := os.Getenv("OPENWEATHER_BASE_URL")
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	fmt.Printf("What city do you want to check the weather: ")
	var city string
	fmt.Scanln(&city)

	request := fmt.Sprintf("%v?q=%v&appid=%v", baseURL, city, apiKey)
	resp, err := http.Get(request)

	if err != nil {
		log.Fatal("Error on request for openweather api:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		byteResponse, _ := io.ReadAll(resp.Body)
		fmt.Println(string(byteResponse))
	}
}
