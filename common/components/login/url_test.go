package login

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	str, err := GetGithubMess("e0a12d92f7cd1d65ee6f78121c15c48d6f9ed7c1")
	fmt.Println(str, err)
}
