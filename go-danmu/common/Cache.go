package common

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"kuukaa.fun/danmu-v5/util"
)

var RedisClient *redis.Client

func InitCache() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: viper.GetString("redis.password"),
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		util.LogError("redis error " + err.Error())
		RedisClient = nil
	} else {
		RedisClient = client
	}
}

func GetCache() *redis.Client {
	return RedisClient
}
