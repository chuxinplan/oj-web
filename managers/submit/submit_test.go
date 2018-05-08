package submit

import (
	"testing"

	"fmt"

	"github.com/open-fightcoder/oj-web/common/g"
)

func TestSubmit(t *testing.T) {
	g.LoadConfig("../../cfg/cfg.toml.debug")
	fmt.Println(IsInOj(4))
}
