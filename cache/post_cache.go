package cache

import "hariprathap-hp/minderaWeatherService/domain/entity"

type Postcache interface {
	Set(key string, value *entity.WeatherInfo)
	Get(key string) *entity.WeatherInfo
}
