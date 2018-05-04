package redis

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-web/common/g"
	"github.com/open-fightcoder/oj-web/common/store"
)

func TestRankGet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitRedis()

	mm, err := PersonMonthRankGet(15)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range mm {
		fmt.Println(v["user_id"])
	}
}
