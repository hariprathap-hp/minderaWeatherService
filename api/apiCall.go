package api

import (
	"encoding/json"
	"fmt"
	"golang-restclient/rest"
	"io/ioutil"
	"minderaWeatherService/config"
	"minderaWeatherService/domain/temp"
	"minderaWeatherService/utils/errors"
	"net/http"
	"time"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://api.weatherstack.com/",
		Timeout: 100 * time.Millisecond,
	}
	weatherstackconn = fmt.Sprintf("current?access_key=%s&query=Melbourne", config.WEATHER_STACK_API_KEY)
	weatherstackURL  = "http://api.weatherstack.com/" + weatherstackconn

	openweatherconn = fmt.Sprintf("weather?q=melbourne,AU&appid=%s", config.OPENWEATHER_API_KEY)
	openWeatherURL  = "http://api.openweathermap.org/data/2.5/" + openweatherconn
)

const ()

func GetWeatherReport() (*temp.WeatherInfo, *errors.RestErr) {
	//create a new request to be sent to the weatherstack api
	req, _ := http.NewRequest("GET", weatherstackURL, nil)
	res, apiErr := http.DefaultClient.Do(req)

	//if the restapi call to weather stack fails, do a rest api call to openweather api
	if apiErr != nil {
		fmt.Println("Openstack not working")
		//As api call to weather stack is not working, an api call is made to openweather api
		result, err := getOpenWeatherReport()
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var tempRes temp.WeatherReport
	//UnMarshall and store the value in the temporary result structure
	json.Unmarshal(body, &tempRes)

	result := temp.WeatherInfo{
		WindSpeed:   float64(tempRes.Current.WindSpeed),
		Temperature: float64(tempRes.Current.Temperature),
	}
	return &result, nil
}

func getOpenWeatherReport() (*temp.WeatherInfo, *errors.RestErr) {
	req, _ := http.NewRequest("GET", openWeatherURL, nil)
	res, apiErr := http.DefaultClient.Do(req)
	if apiErr != nil {
		//if the api call to both the apis fail, then we need to return the value from the storage we have
		return nil, errors.NewInternalServerError("error fetching weather info from both weather stack and open weather apis")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var tempRes temp.OpenWeatherReport
	json.Unmarshal(body, &tempRes)
	result := temp.WeatherInfo{
		WindSpeed:   float64(tempRes.Wind.Speed),
		Temperature: float64(tempRes.Main.Temp),
	}
	return &result, nil
}
