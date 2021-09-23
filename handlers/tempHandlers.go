package handlers

import (
	"minderaWeatherService/domain/temp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TempHandler(c *gin.Context) {
	access_key := c.Request.URL.Query().Get("access_key")
	city := c.Request.URL.Query().Get("query")

	result, tempErr := temp.GetWeatherReport(access_key, city)
	if tempErr != nil {
		c.JSON(tempErr.Status, tempErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
