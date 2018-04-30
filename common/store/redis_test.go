package store

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestList(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	InitRedis()
	isExit := RedisClient.ZScore("person_week_rank", "13")
	fmt.Println(isExit.Err())
}
