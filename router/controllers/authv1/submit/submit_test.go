package submit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"

	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/stretchr/testify/assert"
)

func TestSubmitCommon(t *testing.T) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	resp, err := client.Post("http://127.0.0.1:9001/apiv1/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))

	defer resp.Body.Close()
	if assert.Equal(t, 200, resp.StatusCode, "鉴权失败！") {
		resp, err = client.Post("http://127.0.0.1:9001/authv1/submit/common",
			"application/x-www-form-urlencoded",
			strings.NewReader("problem_id=5&user_id=6&language=JAVA&code=aaaa"))
		if err != nil {
			fmt.Println("POST请求失败: " + err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取Response失败: " + err.Error())
		}
		var respT base.HttpResponse
		if err := json.Unmarshal(body, &respT); err != nil {

			fmt.Println("获取Body失败: " + err.Error())
		}
		assert.Equal(t, 0, respT.Code, "提交失败！")
	}
}

func TestSubmitTest(t *testing.T) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	resp, err := client.Post("http://127.0.0.1:9001/apiv1/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))

	defer resp.Body.Close()
	if assert.Equal(t, 200, resp.StatusCode, "鉴权失败！") {
		resp, err = client.Post("http://127.0.0.1:9001/authv1/submit/test",
			"application/x-www-form-urlencoded",
			strings.NewReader("user_id=1&language=2&input=3&code=4"))
		if err != nil {
			fmt.Println("POST请求失败: " + err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取Response失败: " + err.Error())
		}
		var respT base.HttpResponse
		if err := json.Unmarshal(body, &respT); err != nil {

			fmt.Println("获取Body失败: " + err.Error())
		}
		fmt.Println(string(body))
		assert.Equal(t, 0, respT.Code, "提交失败！")
	}
}

func TestSubmitGetCommon(t *testing.T) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	resp, err := client.Post("http://127.0.0.1:9001/apiv1/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))

	defer resp.Body.Close()
	if assert.Equal(t, 200, resp.StatusCode, "鉴权失败！") {
		resp, err = client.Get("http://127.0.0.1:9001/authv1/submit/getcommon?submit_id=1")
		if err != nil {
			fmt.Println("GET请求失败: " + err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取Response失败: " + err.Error())
		}
		var respT base.HttpResponse
		if err := json.Unmarshal(body, &respT); err != nil {

			fmt.Println("获取Body失败: " + err.Error())
		}
		fmt.Println((string(body)))
		assert.Equal(t, 0, respT.Code, "获取失败！")
	}
}

func TestSubmitGetTest(t *testing.T) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	resp, err := client.Post("http://127.0.0.1:9001/apiv1/login",
		"application/x-www-form-urlencoded",
		strings.NewReader("email=asdfr.com&password=asdfr"))

	defer resp.Body.Close()
	if assert.Equal(t, 200, resp.StatusCode, "鉴权失败！") {
		resp, err = client.Get("http://127.0.0.1:9001/authv1/submit/gettest?submit_id=1")
		if err != nil {
			fmt.Println("GET请求失败: " + err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取Response失败: " + err.Error())
		}
		var respT base.HttpResponse
		if err := json.Unmarshal(body, &respT); err != nil {
			fmt.Println("获取Body失败: " + err.Error())
		}
		fmt.Println(string(body))
		assert.Equal(t, 0, respT.Code, "获取失败！")
	}
}
