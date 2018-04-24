package components

import (
	"fmt"
	"testing"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestList(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	flag := Send("test", &SendMess{"as", 1})
	fmt.Println(flag)
}
