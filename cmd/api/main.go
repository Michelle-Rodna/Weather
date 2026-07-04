package main

import (
	"github.com/labstack/echo/v4"
    "weather/internal/handlers"
	"fmt"
	"github.com/joho/godotenv"
)

func main(){
	if err:= godotenv.Load();err != nil{
		fmt.Println(".env file not found")
	}

	e := echo.New()
	e.GET("/weather", handler.GetHandler)
	e.Start(":8080")
}