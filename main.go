package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type WeatherSuccessResponse struct {
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Humidity int     `json:"humidity"`
	}
	Sys struct {
		Country string `json:"country"`
	}
	Visibility int    `json:"visibility"`
	Name       string `json:"name"`
	Timezone   int    `json:"timezone"`
}

type WeatherNotFoundResponse struct {
	Message string `json:"message"`
}

func main() {
	// https://api.openweathermap.org/data/2.5/weather?q={city name}&appid={API key}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseURL := os.Getenv("OPENWEATHER_BASE_URL")
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	for {
		fmt.Printf("What city do you want to check the weather: ")
		var city string
		fmt.Scanln(&city)

		request := fmt.Sprintf("%v?q=%v&appid=%v", baseURL, city, apiKey)
		response, err := http.Get(request)

		if err != nil {
			log.Fatal("Error on request for openweather api:", err)
		} else {
			handleResponse(response)
			response.Body.Close()
		}

		// Ask if user want to continue
		fmt.Printf("Do you want to search another city? (y/n): ")
		var choice string
		fmt.Scanln(&choice)

		if choice != "y" {
			fmt.Println("Thank you for using Weather CLI!")
			break
		}
	}
}

func handleResponse(response *http.Response) {
	byteResponse, _ := io.ReadAll(response.Body)

	switch response.StatusCode {
	case http.StatusOK:
		weather := WeatherSuccessResponse{}
		json.Unmarshal(byteResponse, &weather)
		// fmt.Printf("API Response: %+v", weather)
		fmt.Printf("The current weather now in %v, %v is %v. The temperature is %.2f°C and the humidity is %v\n",
			weather.Name,
			weather.Sys.Country,
			weather.Weather[0].Description,
			weather.Main.Temp-273.15,
			weather.Main.Humidity,
		)
	case http.StatusNotFound:
		nfWeather := WeatherNotFoundResponse{}
		json.Unmarshal(byteResponse, &nfWeather)
		fmt.Println(strings.Title(nfWeather.Message))
		// fmt.Printf("%v", string(byteResponse))
	default:
		fmt.Printf("Unexpected error occur: %v", response.StatusCode)
	}
	//
	// } else if response.StatusCode == http.StatusNotFound {
	// 	fmt.Println(response)
	// }
}
