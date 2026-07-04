package models

type Response struct{
	Status string `json:"status"`
	Weather *WeatherResponse `json:"weather"`
}

type WeatherResponse struct {
    City      string `json:"city"`
    Temp      int    `json:"temperature"`
    FeelsLike int    `json:"feels_like"`
    Condition string `json:"condition"`
}