package store

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestList(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	client, _ := InitRedis()
	err := client.Set("user", 1, 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := client.Get("user").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
