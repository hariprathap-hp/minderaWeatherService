package handlers

import (
	"minderaWeatherService/domain/temp"
	"minderaWeatherService/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TempHandler(c *gin.Context) {
	access_key := c.Request.URL.Query().Get("access_key")
	city := c.Request.URL.Query().Get("query")

	valErr := validate(access_key, city)
	if valErr != nil {
		c.JSON(valErr.Status, valErr)
		return
	}
	result, tempErr := temp.GetWeatherReport(access_key, city)
	if tempErr != nil {
		c.JSON(tempErr.Status, tempErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func validate(key, city string) *errors.RestErr {
	if key == "" {
		return errors.NewBadRequestError("access key cannot be empty")
	}
	if city == "" {
		return errors.NewBadRequestError("city cannot be empty")
	}

	_, err := strconv.Atoi(city)
	if err == nil {
		return errors.NewBadRequestError("city cannot be an integer value")
	}
	return nil
}
