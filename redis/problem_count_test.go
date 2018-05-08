package redis

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-web/common/g"
	"github.com/open-fightcoder/oj-web/common/store"
)

func TestProblemCountSet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitRedis()

	aa := ProblemCountSet(13, "{\"ac_num\":12,\"total_num\":42}")
	fmt.Println(aa)
}

func TestProblemCountGet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitRedis()

	aa, err := ProblemCountGet(13)
	fmt.Println(aa, err)
}
