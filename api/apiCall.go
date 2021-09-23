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
	weatherstackBaseUrl = "http://api.weatherstack.com/"
	openWeatherBaseUrl  = "http://api.openweathermap.org/data/2.5/"
)

func GetWeatherStackURL(access_key, city string) string {
	weatherstackconn := fmt.Sprintf("current?access_key=%s&query=%s", access_key, city)
	return weatherstackBaseUrl + weatherstackconn
}

func GetOpenWeatherURL() string {
	return openWeatherBaseUrl + fmt.Sprintf("weather?q=melbourne,AU&appid=%s", config.OPENWEATHER_API_KEY)
}
