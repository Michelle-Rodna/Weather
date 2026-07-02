package main
import (
	"fmt"
	"os"
	"weather/internal/weather"
	"github.com/joho/godotenv"
)

func main(){
	if err:= godotenv.Load();err != nil{
		fmt.Println(".env file not found")
	}

	if len(os.Args) < 2{
		fmt.Println("Использование: weather-cil Город")
		fmt.Println("Пример: weather-cil Москва")
		os.Exit(1)
	}
	city := os.Args[1]
	err := weather.GetWeatherByCity(city)
	if err != nil{
		fmt.Println("Ошибка: ", err)
		os.Exit(1)
	}
}