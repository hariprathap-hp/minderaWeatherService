package cacheConfig

import (
	"encoding/json"
	"minderaWeatherService/utils/errors"
	"time"

	"github.com/gregjones/httpcache/diskcache"
	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Second, 5*time.Second)
var diskCache = diskcache.New("./cachedInfo")

type WeatherInfo struct {
	Windspeed   float32 `json:"wind_speed"`
	Temperature float32 `json:"temperature_degrees"`
}

func SetCache(key string, info interface{}) bool {
	Cache.Set(key, info, 3*time.Second)
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

func SetNoExpiredCache(city string, info interface{}) *errors.RestErr {
	bytes, errMarshal := json.Marshal(info)
	if errMarshal != nil {
		return errors.NewInternalServerError("error while marshaling weather info")
	}
	diskCache.Set(city, bytes)
	return nil
}

func GetNoExpiredCache(city string) ([]byte, bool) {
	result, isFound := diskCache.Get(city)
	return result, isFound
}
