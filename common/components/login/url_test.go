package login

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	accessToken, _ := QQCallback("AB6738E6348B957713C98F5DBC973872", "1111")
	openId, _ := GetOpenid(accessToken)
	fmt.Println(openId)
}
