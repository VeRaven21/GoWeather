package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Получение координат города по его названию
func GetGeolocation(ctx context.Context, city string, apikey string) (*Geoposition, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, apikey)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
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
