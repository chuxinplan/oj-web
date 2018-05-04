package login

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	//28AA149D4520BAA0EA7A09879B81A3DE
	//9EC58C000E554465E68F8F51D3D1A1AF
	mess, err := GetQQMess("28AA149D4BAA0EA7A09879B81A3DE", "9EC58C000E554465E68F8F51D3D1A1AF")
	fmt.Println(mess, err.Error())
}
