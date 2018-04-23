package problem

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8000/apiv1/problem/list?origin=1,2,3&tag=三&sort=2&is_asc=1&current_page=1&per_page=10")
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
	assert.Equal(t, 0, respT.Code, "获取失败！")
}

func TestGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8000/apiv1/problem/get?id=1")
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
	assert.Equal(t, 0, respT.Code, "获取失败！")
}

func TestRandom(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8000/apiv1/problem/random?origin=1,2,3&tag=三")
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
	assert.Equal(t, 0, respT.Code, "获取失败！")
}
