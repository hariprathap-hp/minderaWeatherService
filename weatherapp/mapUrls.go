package weatherapp

import (
	"Melbourne_Weather_Service/api"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mtemp struct {
	WindSpeed    int `json:"wind_speed"`
	Temperrature int `json:"temperature_degrees"`
}

func mapUrl() {
	fmt.Println("Map Urls")
	router.GET("/v1/weather/", tempHandler)
}

func tempHandler(c *gin.Context) {
	result, tempErr := api.GetWeatherReport()
	if tempErr != nil {
		c.JSON(tempErr.Status, tempErr)
		return
	}
	c.JSON(http.StatusOK, result)
	//c.JSON(http.StatusOK, result)
}
