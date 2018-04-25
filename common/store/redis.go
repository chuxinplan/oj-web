package store

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var client *redis.Client
var once sync.Once

func InitRedis() *redis.Client {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		pong, err := client.Ping().Result()
		fmt.Println(pong, err)
	})
	return client
}

func CloseRedis() {

}
