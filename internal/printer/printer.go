package printer

import (
	"fmt"
	"math"

	"github.com/VeRaven21/GoWeather/internal/api"
)

// ANSI цвета
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// Направление ветра по градусам
func windDirection(deg float64) string {
	directions := []string{"↑ N", "↗ NE", "→ E", "↘ SE", "↓ S", "↙ SW", "← W", "↖ NW"}
	index := int(math.Mod(deg+22.5, 360) / 45)
	return directions[index]
}

// Форматирование температуры
func tempColor(temp float64) string {
	if temp > 25 {
		return Red + fmt.Sprintf("%.1f°C", temp) + Reset
	} else if temp > 10 {
		return Yellow + fmt.Sprintf("%.1f°C", temp) + Reset
	} else if temp > 0 {
		return Green + fmt.Sprintf("%.1f°C", temp) + Reset
	} else {
		return Cyan + fmt.Sprintf("%.1f°C", temp) + Reset
	}
}

func PrintWeather(w api.Weather) {
	fmt.Println(White + "📍 Текущая погода" + Reset)
	fmt.Printf("🌡️  Температура: %s\n", tempColor(w.CurrentWeather.Temperature))
	fmt.Printf("💧 Влажность:   %s%.0f%%\n", Blue, w.CurrentWeather.Humidity)
	fmt.Printf("💨 Ветер:       %.1f м/с %s (%s)\n", w.CurrentWeather.WindSpeed, windDirection(w.CurrentWeather.WindDirection), Blue+fmt.Sprintf("%.0f°", w.CurrentWeather.WindDirection)+Reset)

	// Осадки
	precip := w.CurrentWeather.Rain + w.CurrentWeather.Showers + w.CurrentWeather.Snowfall
	if precip > 0 {
		if w.CurrentWeather.Snowfall > 0 {
			fmt.Printf("❄️  Снег:        %.1f мм\n", w.CurrentWeather.Snowfall)
		} else if w.CurrentWeather.Showers > 0 {
			fmt.Printf("🚿 Ливень:      %.1f мм\n", w.CurrentWeather.Showers)
		} else {
			fmt.Printf("🌧️  Дождь:       %.1f мм\n", w.CurrentWeather.Rain)
		}
	}

	fmt.Println("\n" + White + "📅 Прогноз на ближайшие дни" + Reset)
	for i, date := range w.Daily.Time {
		if i >= 5 { // Показываем только 5 дней
			break
		}
		maxT := tempColor(w.Daily.TempMax[i])
		minT := tempColor(w.Daily.TempMin[i])
		rain := w.Daily.RainSum[i]

		// Иконка по осадкам и температуре
		var icon string
		if rain > 0 {
			icon = "🌧️"
		} else if w.Daily.TempMax[i] > 20 {
			icon = "☀️"
		} else {
			icon = "☁️"
		}

		rainStr := ""
		if rain > 0 {
			rainStr = fmt.Sprintf(" 🌧️%.1fмм", rain)
		}

		fmt.Printf("%s %s: %s / %s%s\n", icon, date, maxT, minT, rainStr)
	}
	fmt.Println()
}
