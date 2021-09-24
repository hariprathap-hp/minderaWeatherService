package handlers

import (
	"minderaWeatherService/domain/temp"
	"minderaWeatherService/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TempHandler(c *gin.Context) {
	city := c.Request.URL.Query().Get("city")

	valErr := validate(city)
	if valErr != nil {
		c.JSON(valErr.Status, valErr)
		return
	}
	result, tempErr := temp.GetWeatherReport(city)
	if tempErr != nil {
		c.JSON(tempErr.Status, tempErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func validate(city string) *errors.RestErr {
	if city == "" {
		return errors.NewBadRequestError("city cannot be empty")
	}

	_, err := strconv.Atoi(city)
	if err == nil {
		return errors.NewBadRequestError("city cannot be an integer value")
	}
	return nil
}
