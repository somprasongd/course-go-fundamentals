package main

import "fmt"

// WeatherResponse คือโครงสร้างที่ API ของเราต้องการ
type WeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
}

// ExternalWeatherService จำลองบริการภายนอกที่ให้ข้อมูลในรูปแบบ map
type ExternalWeatherService struct{}

func (e *ExternalWeatherService) GetWeather(location string) map[string]interface{} {
	// จำลองการเรียก API ภายนอก
	return map[string]interface{}{
		"temp": 30.5,
		"desc": "Sunny",
	}
}

// WeatherServiceAdapter ปรับข้อมูลจาก ExternalWeatherService ให้เป็น WeatherResponse
type WeatherServiceAdapter struct {
	external *ExternalWeatherService
}

func (a *WeatherServiceAdapter) GetWeather(location string) WeatherResponse {
	data := a.external.GetWeather(location)
	return WeatherResponse{
		Temperature: data["temp"].(float64),
		Condition:   data["desc"].(string),
	}
}

func main() {
	external := &ExternalWeatherService{}
	adapter := &WeatherServiceAdapter{external: external}
	weather := adapter.GetWeather("Phuket")

	fmt.Printf("Temperature: %.2f\n", weather.Temperature)
	fmt.Printf("Condition: %s\n", weather.Condition)
}
