package store

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/open-fightcoder/oj-web/common/g"
)

var client *redis.Client
var once sync.Once

func InitRedis() (*redis.Client, error) {
	once.Do(func() {
		cfg := g.Conf().Redis
		client = redis.NewClient(&redis.Options{
			Addr:     cfg.Address,
			Password: cfg.Password,
			DB:       cfg.Database,
		})
	})
	_, err := client.Ping().Result()
	if err != nil {
		//write log
		return nil, err
	}
	return client, nil
}

func CloseRedis() {

}
