package weatherapp

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	fmt.Println("Weather App")
	mapUrl()
	router.Run(":8080")
}
