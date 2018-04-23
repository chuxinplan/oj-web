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

//func TestCodeGet(t *testing.T) {
//	var client http.Client
//	jar, err := cookiejar.New(nil)
//	if err != nil {
//		panic(err)
//	}
//	client.Jar = jar
//
//	resp, err := client.Post("http://127.0.0.1:8000/apiv1/login",
//		"application/x-www-form-urlencoded",
//		strings.NewReader("email=asdfr.com&password=asdfr"))
//
//	defer resp.Body.Close()
//	_, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("读取Response失败: " + err.Error())
//	}
//
//	fmt.Println(resp.Header)
//
//	resp, err = client.Get("http://127.0.0.1:8000/authv1/problem/code/get?user_id=6&problem_id=5")
//
//	if err != nil {
//		fmt.Println("GET请求失败: " + err.Error())
//	}
//	defer resp.Body.Close()
//	if assert.Equal(t, 200, resp.StatusCode, "鉴权失败！") {
//
//		body, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			fmt.Println("读取Response失败: " + err.Error())
//		}
//		fmt.Println(string(body))
//		var respT base.HttpResponse
//		if err := json.Unmarshal(body, &respT); err != nil {
//
//			fmt.Println("获取Body失败: " + err.Error())
//		}
//		fmt.Println((string(body)))
//		assert.Equal(t, 0, respT.Code, "获取代码失败！")
//	}
//}

func TestCodeGet(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8000/authv1/problem/code/get?user_id=1&problem_id=2")
	if err != nil {
		fmt.Println("Get请求失败: " + err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取Response失败: " + err.Error())
	}
	fmt.Println(string(body))
	var respT base.HttpResponse
	if err := json.Unmarshal(body, &respT); err != nil {
		fmt.Println("获取Body失败: " + err.Error())
	}
	assert.Equal(t, 0, respT.Code, "获取失败！")
}
