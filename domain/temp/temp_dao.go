package temp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"minderaWeatherService/api"
	"minderaWeatherService/clients/rest_client"
	"minderaWeatherService/config/cacheConfig"
	"minderaWeatherService/utils/errors"
)

func parseCacheResult(cacheInfo interface{}) *WeatherInfo {
	cachedRes, _ := json.Marshal(cacheInfo)
	var result WeatherInfo
	json.Unmarshal(cachedRes, &result)
	return &result
}

func GetWeatherReport(city string) (*WeatherInfo, *errors.RestErr) {
	//check if the weather info is present in the cache. If present in cache, return the result as json
	cacheInfo, isCached := cacheConfig.GetCache(city)
	if isCached {
		fmt.Println("info cached")
		cachedRes := parseCacheResult(cacheInfo)
		return cachedRes, nil
	}
	fmt.Println("info not cached")
	//create a new request to be sent to the weatherstack api if weather info not present in cache
	res, apiErr := rest_client.Get(api.GetWeatherStackURL(city))

	//if the restapi call to weatherstack fails, do a rest api call to openweather api
	if apiErr != nil {
		result, err := getOpenWeatherReport(city)
		//if weather info is not returned from openweather api as well, check cache if result is present
		if err != nil {
			cacheInfo, isCached := cacheConfig.GetCache(city)
			if isCached {
				cachedRes := parseCacheResult(cacheInfo)
				return cachedRes, nil
			} else {
				//if not present in temporary cache, check permanent file for weather info
				return nil, nil
			}
		}
		//set the result from openweather api in the cache
		cacheConfig.SetCache(city, result)
		//set permanent cache
		return result, nil
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var tempRes WeatherReport
	json.Unmarshal(body, &tempRes)
	result := WeatherInfo{
		WindSpeed:   float64(tempRes.Current.WindSpeed),
		Temperature: float64(tempRes.Current.Temperature),
	}

	//set the result from weatherstack api in the cache
	cacheConfig.SetCache(city, result)
	return &result, nil
}

func getOpenWeatherReport(city string) (*WeatherInfo, *errors.RestErr) {
	res, apiErr := rest_client.Get(api.GetOpenWeatherURL(city))
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
