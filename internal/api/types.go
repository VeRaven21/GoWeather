package api

type GeopositionResponse struct {
	Results []Geoposition `json:"results"`
}

type Geoposition struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

type CurrentWeather struct {
	Temperature   float64 `json:"temperature_2m"`
	Rain          float64 `json:"rain"`
	Humidity      float64 `json:"relative_humidity_2m"`
	Showers       float64 `json:"showers"`
	Snowfall      float64 `json:"snowfall"`
	WindSpeed     float64 `json:"wind_speed_10m"`
	WindDirection float64 `json:"wind_direction_10m"`
}

type DailyWeather struct {
	Time        []string  `json:"time"`
	WeatherCode []int     `json:"weathercode"`
	TempMax     []float64 `json:"temperature_2m_max"`
	TempMin     []float64 `json:"temperature_2m_min"`
	RainSum     []float64 `json:"rain_sum"`
}

type Weather struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`

	CurrentWeather CurrentWeather `json:"current"`
	Daily          DailyWeather   `json:"daily"`
}
