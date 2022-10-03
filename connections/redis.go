package connections

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func RedisConn(host string, port int) (redisClient *redis.Client, err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("redis conneted!")
	return redisClient, nil
}
