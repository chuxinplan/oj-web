package problem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/stretchr/testify/assert"
)

func TestCollectionGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9001/authv1/problem/collection/get?problem_ids=1,2,3,4")
	if err != nil {
		fmt.Println("GET请求失败: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取Response失败: " + err.Error())
	}
	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println("获取Body失败: " + err.Error())
	}
	assert.Equal(t, 0, respT.Code, "收藏失败！")
}

func TestCollectionSet(t *testing.T) {
	resp, err := http.Post("http://127.0.0.1:9001/authv1/problem/collection/set",
		"application/x-www-form-urlencoded",
		strings.NewReader("problem_id=1&flag=set"))
	if err != nil {
		fmt.Println("POST请求失败: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取Response失败: " + err.Error())
	}

	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println("获取Body失败: " + err.Error())
	}
	assert.Equal(t, 0, respT.Code, "收藏失败！")
}
