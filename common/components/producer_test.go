package components

import (
	"fmt"
	"testing"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestList(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	flag := Send("vjudger", &SendMess{"", 3})
	fmt.Println(flag)
}
