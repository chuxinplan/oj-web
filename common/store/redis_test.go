package store

import (
	"fmt"
	"testing"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestList(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	fmt.Println(g.Conf().Redis.Address, g.Conf().Redis.Password, g.Conf().Redis.Database)
}
