package weatherapp

import (
	"minderaWeatherService/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapUrl() {
	router.GET("/v1/weather/", tempHandler)
}

func tempHandler(c *gin.Context) {
	result, tempErr := api.GetWeatherReport()
	if tempErr != nil {
		c.JSON(tempErr.Status, tempErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
