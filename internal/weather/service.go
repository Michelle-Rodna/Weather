package weather
import (
	"fmt"
	"weather/internal/client"
	"weather/geocoder"
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