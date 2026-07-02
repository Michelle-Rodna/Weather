package client
import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
	"os"
)

type WeatherResponse struct{
	Now int `json:"now"`
	Fact Fact `json:"fact"`
}

type Fact struct{
	Temp int `json:"temp"`
	FeelsLike int `json:"feels_like"`
	Condition string `json:"condition"`
}

func GetWeather(lat, lon string)(*WeatherResponse, error){

	apiKey := os.Getenv("YANDEX_WEATHER_API_KEY")
	if apiKey == ""{
		return nil, fmt.Errorf("API key not found. Please set it.")
	}

	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?lat=%s&lon=%s", lat, lon)

	req, err := http.NewRequest("GET", url,  nil)
	if err != nil{
		return nil, err
	}
	req.Header.Add("X-Yandex-Weather-Key", apiKey)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Api return the %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}