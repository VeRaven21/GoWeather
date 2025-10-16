package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Получение координат города по его названию
func GetGeolocation(city string) (*Geoposition, error) {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", city)

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

	var geopos GeopositionResponse
	err = json.Unmarshal(respData, &geopos)
	if err != nil {
		return nil, err
	}

	if len(geopos.Results) == 0 {
		return nil, fmt.Errorf("city not found")
	}

	return &geopos.Results[0], nil
}

func GetWeather(location Geoposition) (*Weather, error) {
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
