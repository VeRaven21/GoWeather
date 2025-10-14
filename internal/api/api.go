package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Получение координат города по его названию
func GetGeolocation(city string, api_key string) (*Geoposition, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, api_key)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geopos []Geoposition
	err = json.Unmarshal(respData, &geopos)
	if err != nil {
		return nil, err
	}

	return &geopos[0], nil
}

func GetWeather(location Geoposition, api_key string) (*Weather, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&daily=weather_code,temperature_2m_max,temperature_2m_min,rain_sum&current=temperature_2m,rain,weather_code,relative_humidity_2m,showers,snowfall,surface_pressure,wind_speed_10m,wind_direction_10m&timezone=auto", location.Lat, location.Lon)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var weather Weather
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, fmt.Errorf("failed to decode weather data: %w", err)
	}

	return &weather, nil
}
