# GoWeather

Console application to get weather from terminal

# Usage
~~Add your Openweather api key in `./cmd/main.go` file and run ~~

Create file named `.env` and past following line
```
OPENWEATHER_API_KEY = ...
```
with your api key, then run 
```bash
go run ./cmd/main.go <city name>
```

Writes weather for today and next 5 days 
```
📍 Текущая погода
🌡️  Температура: 4.2°C
💧 Влажность:   59%
💨 Ветер:       16.9 м/с ↘ SE (151°)

📅 Прогноз на ближайшие дни
🌧️ 2025-10-14: 5.3°C / -0.1°C 🌧️0.9мм
🌧️ 2025-10-15: 3.4°C / 2.0°C 🌧️5.9мм
🌧️ 2025-10-16: 5.5°C / 2.0°C 🌧️2.4мм
🌧️ 2025-10-17: 3.7°C / 1.0°C 🌧️1.5мм
☁️ 2025-10-18: 5.0°C / 0.4°C
```

# To do
- [x] Add actual weather forecast 
- [x] Add better visuals

- [ ] ~~Add configuration of api key and default city~~ City location Api that doesn't requre api key

- [ ] Different languages output


