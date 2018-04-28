package store

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestList(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	InitRedis()
	isExit := RedisClient.ZLexCount("person_week_rank", "[6", "[6")
	fmt.Println(isExit.Val())
}
