package api

import (
	"fmt"
	"minderaWeatherService/config/keyConfig"
)

const (
	weatherstackBaseUrl = "http://api.weatherstack.com/"
	openWeatherBaseUrl  = "http://api.openweathermap.org/data/2.5/"
)

func GetWeatherStackURL(city string) string {
	weatherstackconn := fmt.Sprintf("current?access_key=%s&query=%s", keyConfig.WEATHER_STACK_API_KEY, city)
	return weatherstackBaseUrl + weatherstackconn
}

func GetOpenWeatherURL(city string) string {
	openweatherconn := fmt.Sprintf("weather?q=%s,AU&appid=%s", city, keyConfig.OPENWEATHER_API_KEY)
	return openWeatherBaseUrl + openweatherconn
}
