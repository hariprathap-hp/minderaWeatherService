package temp

import (
	"encoding/json"
	"fmt"
	"hariprathap-hp/minderaWeatherService/api"
	"hariprathap-hp/minderaWeatherService/cache"
	"hariprathap-hp/minderaWeatherService/clients/rest_client"
	"hariprathap-hp/minderaWeatherService/domain/entity"
	"hariprathap-hp/minderaWeatherService/utils/errors"
	"io/ioutil"

	"strings"
)

func parseCacheResult(cacheInfo interface{}) *entity.WeatherInfo {
	cachedRes, _ := json.Marshal(cacheInfo)
	var result entity.WeatherInfo
	json.Unmarshal(cachedRes, &result)
	return &result
}

var postcache cache.Postcache = cache.NewRedisCache("localhost:6379", 1, 10)

func GetWeatherReport(city string) (*entity.WeatherInfo, *errors.RestErr) {
	fmt.Println("GetWeatherReport")
	//check if the weather info is present in the cache. If present in cache, return the result as json
	var weather *entity.WeatherInfo = postcache.Get(city)
	if weather != nil {
		cachedRes := parseCacheResult(weather)
		return cachedRes, nil
	}
	//create a new request to be sent to the weatherstack api if weather info not present in cache
	res, apiErr := rest_client.Get(api.GetWeatherStackURL(city))

	//if the restapi call to weatherstack fails, do a rest api call to openweather api
	if apiErr != nil {
		result, err := getOpenWeatherReport(city)
		//if weather info is not returned from openweather api as well, check cache if result is present
		if err != nil {
			var weather *entity.WeatherInfo = postcache.Get(city)
			if weather != nil {
				cachedRes := parseCacheResult(weather)
				return cachedRes, nil
			}
		}
		//set the result from openweather api in the cache
		postcache.Set(city, result)
		return result, nil
	}
	//If no error is returned by openstack weather api, proceed storing in cache and returning the result
	body, _ := ioutil.ReadAll(res.Body)
	//if no error is returned, but returns 615 due to wrong city name
	if strings.Contains(string(body), "615") {
		return nil, errors.NewNotFoundError("weather info not found. check the city name entered")
	}
	defer res.Body.Close()

	var tempRes entity.WeatherReport
	json.Unmarshal(body, &tempRes)
	result := entity.WeatherInfo{
		WindSpeed:   float64(tempRes.Current.WindSpeed),
		Temperature: float64(tempRes.Current.Temperature),
	}

	//set the result from weatherstack api in the cache
	postcache.Set(city, &result)

	return &result, nil
}

func getOpenWeatherReport(city string) (*entity.WeatherInfo, *errors.RestErr) {
	res, apiErr := rest_client.Get(api.GetOpenWeatherURL(city))
	if apiErr != nil {
		//if the api call to both the apis fail, then we need to return the value from the storage we have
		return nil, errors.NewInternalServerError("error fetching weather info from both weather stack and open weather apis")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if strings.Contains(string(body), "404") {
		return nil, errors.NewNotFoundError("weather info not found. check the city name entered")
	}
	var tempRes entity.OpenWeatherReport
	json.Unmarshal(body, &tempRes)
	result := entity.WeatherInfo{
		WindSpeed:   float64(tempRes.Wind.Speed),
		Temperature: float64(tempRes.Main.Temp),
	}
	result.FTOC()
	return &result, nil
}
