package weatherapp

import "minderaWeatherService/handlers"

func mapUrl() {
	router.GET("/v1/weather/", handlers.TempHandler)
}
