package location

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)

type GeoData struct {
	City string `json:"city"`
}

func GetCityByIp(ip string)(city string, err error){
	url := fmt.Sprintf("http://ip-api.com/json/%s",ip)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return "", fmt.Errorf("Api return the %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}
	var data GeoData
	err = json.Unmarshal(body, &data)
	if err != nil{
		return "", err
	}
	return data.City, nil
}
