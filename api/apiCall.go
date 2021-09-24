package api

import (
	"fmt"
	"golang-restclient/rest"
	"minderaWeatherService/config"
	"time"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://api.weatherstack.com/",
		Timeout: 100 * time.Millisecond,
	}
)

const (
	weatherstackBaseUrl = "http://api.weatherstack11.com/"
	openWeatherBaseUrl  = "http://api.openweathermap.org/data/2.5/"
)

func GetWeatherStackURL(city string) string {
	weatherstackconn := fmt.Sprintf("current?access_key=%s&query=%s", config.WEATHER_STACK_API_KEY, city)
	return weatherstackBaseUrl + weatherstackconn
}

func GetOpenWeatherURL(city string) string {
	return openWeatherBaseUrl + fmt.Sprintf("weather?q=%s,AU&appid=%s", city, config.OPENWEATHER_API_KEY)
}
