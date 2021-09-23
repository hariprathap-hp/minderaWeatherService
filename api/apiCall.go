package api

import (
	"Melbourne_Weather_Service/config"
	"Melbourne_Weather_Service/domain/temp"
	"Melbourne_Weather_Service/utils/errors"
	"encoding/json"
	"fmt"
	"golang-restclient/rest"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://api.weatherstack.com/",
		Timeout: 100 * time.Millisecond,
	}
	connString = fmt.Sprintf("current?access_key=%s&query=Melbourne", config.WEATHER_STACK_API_KEY)
	url        = "http://api.weatherstack.com/" + connString
)

const ()

func GetWeatherReport() (*temp.WeatherReport, *errors.RestErr) {
	req, _ := http.NewRequest("GET", url, nil)
	res, apiErr := http.DefaultClient.Do(req)
	if apiErr != nil {
		return nil, errors.NewInternalServerError("api call to weatherstack api failed")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//response := oauthRestClient.Get(connString)
	var result temp.WeatherReport
	json.Unmarshal(body, &result)
	return &result, nil
}
