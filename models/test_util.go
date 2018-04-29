package models

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/open-fightcoder/oj-web/common"
)

var once sync.Once

func InitAllInTest() {
	once.Do(func() {
		common.Init("../cfg/cfg.toml.debug")
	})
}
