package api

import (
	"fmt"
	"minderaWeatherService/config/keyConfig"
	"strings"
)

const (
	weatherstackBaseUrl      = "http://api.weatherstack111.com/"
	openWeatherBaseUrl       = "http://api.openweathermap.org/data/2.5/"
	falseweatherstackBaseUrl = "http://api.weatherstack11.com/"
	falseopenWeatherBaseUrl  = "http://api.openweathermap11.org/data/2.5/"
)

func GetWeatherStackURL(city string) string {
	weatherstackconn := fmt.Sprintf("current?access_key=%s&query=%s", keyConfig.WEATHER_STACK_API_KEY, city)
	if strings.Compare(city, "canberra") == 0 {
		return falseweatherstackBaseUrl + weatherstackconn
	}
	return weatherstackBaseUrl + weatherstackconn
}

func GetOpenWeatherURL(city string) string {
	openweatherconn := fmt.Sprintf("weather?q=%s,AU&appid=%s", city, keyConfig.OPENWEATHER_API_KEY)
	if strings.Compare(city, "canberra") == 0 {
		return falseopenWeatherBaseUrl + openweatherconn
	}
	return openWeatherBaseUrl + openweatherconn
}
