package temp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"minderaWeatherService/api"
	"minderaWeatherService/utils/errors"
	"net/http"
)

func GetWeatherReport(access_key, city string) (*WeatherInfo, *errors.RestErr) {
	//create a new request to be sent to the weatherstack api
	req, _ := http.NewRequest("GET", api.GetWeatherStackURL(access_key, city), nil)
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

	var tempRes WeatherReport
	//UnMarshall and store the value in the temporary result structure
	json.Unmarshal(body, &tempRes)

	result := WeatherInfo{
		WindSpeed:   float64(tempRes.Current.WindSpeed),
		Temperature: float64(tempRes.Current.Temperature),
	}
	return &result, nil
}

func getOpenWeatherReport() (*WeatherInfo, *errors.RestErr) {
	fmt.Println(api.GetOpenWeatherURL())
	req, _ := http.NewRequest("GET", api.GetOpenWeatherURL(), nil)
	res, apiErr := http.DefaultClient.Do(req)
	if apiErr != nil {
		//if the api call to both the apis fail, then we need to return the value from the storage we have
		return nil, errors.NewInternalServerError("error fetching weather info from both weather stack and open weather apis")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var tempRes OpenWeatherReport
	json.Unmarshal(body, &tempRes)
	result := WeatherInfo{
		WindSpeed:   float64(tempRes.Wind.Speed),
		Temperature: float64(tempRes.Main.Temp),
	}
	result.FTOC()
	return &result, nil
}