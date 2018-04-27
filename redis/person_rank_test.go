package redis

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestRankGet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")

	strs, err := PersonWeekRankGet(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(strs)
}
