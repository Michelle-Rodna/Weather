package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"weather/internal/models"
	"weather/internal/weather"
	"weather/internal/location"
)

func GetHandler(c echo.Context)error{
	city := c.QueryParam("city")
	if city == ""{
		userId := c.RealIP()
		if userId == "127.0.0.1" || userId == "::1"{
			return c.JSON(http.StatusBadRequest, models.Response{
				Status: "Error",
				Weather: nil,
			})
		}
		var err error
		city, err = location.GetCityByIp(userId)
		if err != nil{
			return c.JSON(http.StatusBadRequest, models.Response{
				Status: "Error",
				Weather: nil,
			})
		}
	}
	info, err := weather.GetWeatherByCityServer(city)
	if err != nil{
		return c.JSON(http.StatusBadRequest, models.Response{
			Status: "Error",
			Weather: nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status: "Success",
		Weather: info,
	})
}