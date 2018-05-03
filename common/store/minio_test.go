package store

import (
	"testing"

	"github.com/open-fightcoder/oj-web/common/g"
	. "github.com/open-fightcoder/oj-web/managers"
)

func TestMinio(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	InitMinio()
	err := UpdateCode("1525336932.txt", "aaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		panic(err)
	}
}
