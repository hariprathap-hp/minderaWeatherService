package cache

import (
	"context"
	"encoding/json"
	"hariprathap-hp/minderaWeatherService/domain/entity"
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) Postcache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *entity.WeatherInfo) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(context.TODO(), key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *entity.WeatherInfo {
	client := cache.getClient()
	val, err := client.Get(context.TODO(), key).Result()
	if err != nil {
		return nil
	}
	weather := entity.WeatherInfo{}
	json.Unmarshal([]byte(val), &weather)
	return &weather
}
