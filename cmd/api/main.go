package main

import (
	"log"
    "weather/internal/cache"
	"github.com/labstack/echo/v4"
    "weather/internal/handlers"
	"fmt"
	"github.com/joho/godotenv"
)

func main(){
	if err:= godotenv.Load();err != nil{
		fmt.Println(".env file not found")
	}

	if err := cache.InitRedis("localhost:6379", "", 0); err != nil {
        log.Fatal("Redis init error:", err)
    }
    log.Println("Redis connected")

	e := echo.New()
	e.GET("/weather", handler.GetHandler)
	e.Start(":8080")
}