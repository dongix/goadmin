package connection

import (
	"config"
	"fmt"
	"log"

	"exception"

	"github.com/go-redis/redis"
)

var options redis.Options

func init() {
	conf := new(config.RedisConfig)
	error := config.GetConfig(conf)
	if error != nil {
		log.Fatal("read redis config error!")
		panic(error)
	}
	options = redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	}
}

// GetRedisClient 获取redis 客户端
func GetRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&options)
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("create redis client error:")
		log.Fatal(err.Error())
		return client, err
	}
	if pong != "PONG" {
		return client, &exception.ErrorCode{
			Code: "-1",
			Msg:  "redis connect error",
		}
	}
	return client, nil
}
