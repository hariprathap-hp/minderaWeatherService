package cacheConfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"minderaWeatherService/utils/errors"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Second, 5*time.Second)

type WeatherInfo struct {
	Windspeed   float32 `json:"wind_speed"`
	Temperature float32 `json:"temperature_degrees"`
}

func SetCache(key string, info interface{}) bool {
	Cache.Set(key, info, 10*time.Second)
	return true
}

func GetCache(key string) (interface{}, bool) {
	var found bool
	data, found := Cache.Get(key)
	if !found {
		return nil, false
	}
	return data, found
}

func SetNoExpiredCache(info interface{}) (*errors.RestErr, bool) {
	filecache, err := os.Create("weather_info.json")
	if err != nil {
		fmt.Println("file opening err")
	}
	res, errMarshal := json.Marshal(info)
	if errMarshal != nil {
		return errors.NewInternalServerError("error while json encoding of the input received"), false
	}
	_, errWrite := filecache.Write(res)
	if errWrite != nil {
		return errors.NewInternalServerError("error while writing temperature info in file"), false
	}
	return nil, true
}

func GetNoExpiredCache() (*WeatherInfo, *errors.RestErr) {
	filecache, errOpen := os.Open("weather_info.json")
	if errOpen != nil {
		return nil, errors.NewInternalServerError("error while opening file to fetch weather info")
	}
	res, errRead := ioutil.ReadAll(filecache)
	if errRead != nil {
		return nil, errors.NewInternalServerError("error while reading data from file reader")
	}
	var result WeatherInfo
	errUnMarshal := json.Unmarshal(res, &result)
	if errUnMarshal != nil {
		return nil, errors.NewInternalServerError("error while unmarshaling weather info read from file")
	}
	return &result, nil
}
