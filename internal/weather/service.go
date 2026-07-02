package weather
import (
	"fmt"
	"weather/internal/client"
)

func GetWeatherByCity(city string)error{
	if city != "Москва"{
		return fmt.Errorf("Город %s не поддерживается.", city)
	}

	info, err := client.GetWeather("55.7558", "37.6173")
	if err != nil{
		return err
	}

	fmt.Printf("Температура: %d\n", info.Fact.Temp)
	fmt.Printf("Ощущается: %d\n", info.Fact.FeelsLike)
	fmt.Printf("Состояние: %s\n", info.Fact.Condition)

	return nil
}