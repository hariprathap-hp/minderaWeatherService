package main

import (
	"minderaWeatherService/handlers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/v1/weather/", handlers.TempHandler)
	router.Run(":8080")
}
