package apiv1

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSomething(t *testing.T) {
	resp, err := http.Post("http://127.0.0.1:8000/apiv1/account/register",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=test.com&password=aaa"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &mapResult); err != nil {
		fmt.Println(err)
	}
	code := mapResult["code"].(float64)
	assert.Equal(t, 0, int(code), "注册失败！")
}
