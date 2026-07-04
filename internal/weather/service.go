package weather
import (
	"fmt"
	"weather/internal/client"
	"weather/internal/geocoder"
	"weather/internal/models"
)

func GetWeatherByCity(city string)error{
	lat, lon, err := geocoder.GetCoordinates(city)
	if err != nil{
		return fmt.Errorf("Error: %w", err)
	}
	info, err := client.GetWeather(lat, lon)
	if err != nil{
		return err
	}

	fmt.Printf("Температура: %d\n", info.Fact.Temp)
	fmt.Printf("Ощущается: %d\n", info.Fact.FeelsLike)
	fmt.Printf("Состояние: %s\n", info.Fact.Condition)

	return nil
}

func GetWeatherByCityServer(city string)(*models.WeatherResponse, error){
	lat, lon, err := geocoder.GetCoordinates(city)
	if err != nil{
		return nil, fmt.Errorf("Error: %w", err)
	}
	info, err := client.GetWeather(lat, lon)
	if err != nil{
		return nil, err
	}

	return &models.WeatherResponse{
		City:      city,
        Temp:      info.Fact.Temp,
        FeelsLike: info.Fact.FeelsLike,
        Condition: info.Fact.Condition,
	}, nil
}