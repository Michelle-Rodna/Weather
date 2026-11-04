package grpcserver
import (
	"context"
	"weather/weatherpb"
	"weather/internal/weather"
)

type Server struct{
	weatherpb.UnimplementedWeatherServiceServer
}

func (s *Server) GetWeather (ctx context.Context, req *weatherpb.WeatherRequest) (*weatherpb.WeatherResponse, error){
	if err := ctx.Err(); err != nil{
		return nil, err
	}
	
	info, err := weather.GetWeatherByCityServer(req.City)
	if err != nil{
		return nil, err
	}
	return &weatherpb.WeatherResponse{
		City: info.City,
		Temp: int32(info.Temp),
		FeelsLike: int32(info.FeelsLike),
		Condition: info.Condition,
	}, nil
}