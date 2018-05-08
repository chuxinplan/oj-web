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

	aa := ProblemCountSet(13, "accepted\":1,\"fail_num\":3,\"wrong_answer\":1,\"compilation_error\":1,\"time_limit_exceeded\":5,\"memory_limit_exceeded\":1,\"output_limit_exceeded\":1,\"runtime_error\":3,\"system_error\":2}")
	fmt.Println(aa)
}

func TestProblemCountGet(t *testing.T) {
	g.LoadConfig("../cfg/cfg.toml.debug")
	store.InitRedis()

	aa, err := ProblemCountGet(13)
	fmt.Println(aa, err)
}
