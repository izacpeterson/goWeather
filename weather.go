package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
	WeatherOverview string `json:"weather_overview"`
}

func main() {
	url := "https://api.openweathermap.org/data/3.0/onecall/overview?lat=40.239&lon=-111.656&appid=35701ad524635b065706933439a9ab30&units=imperial"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var weatherResponse WeatherResponse

	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		fmt.Println("Error unmarshalling: %v", err)
		return
	}

	fmt.Printf("\nCurrent weather in Provo, Utah\n%v\n", weatherResponse.WeatherOverview)

}
