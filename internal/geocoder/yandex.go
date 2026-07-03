package geocoder

import (
	"net/http"
	"fmt"
	"io"
	"encoding/json"
	"os"
	"strings"
	"net/url"
)

type GeocodeResponse struct {
    Response struct {
        GeoObjectCollection struct {
            FeatureMember []struct {
                GeoObject struct {
                    Point struct {
                        Pos string `json:"pos"`
                    } `json:"Point"`
                } `json:"GeoObject"`
            } `json:"featureMember"`
        } `json:"GeoObjectCollection"`
    } `json:"response"`
}

func GetCoordinates(city string)(lat, lon string, err error){
	apiKey := os.Getenv("YANDEX_GEO_API_KEY")
	if apiKey == ""{
		return "", "", fmt.Errorf("API key not found. Please set it.")
	}

	url := fmt.Sprintf("https://geocode-maps.yandex.ru/v1/?apikey=%s&geocode=%s&format=json", apiKey, url.QueryEscape(city))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return "", "", err 
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return "", "", fmt.Errorf("Error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return "", "", err
	}

	var geo GeocodeResponse
	err = json.Unmarshal(body, &geo)
	if err != nil{
		return "", "", err
	}
	
	if len(geo.Response.GeoObjectCollection.FeatureMember) == 0{
		return "", "", fmt.Errorf("City not found: %s", city)
	}
	pos := geo.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Point.Pos
	coord := strings.Split(pos," ")
	if len(coord) != 2{
		return "", "", fmt.Errorf("Got not right data: %s", pos)
	}
	lat = coord[1]
	lon = coord[0]

	return lat, lon, nil
}