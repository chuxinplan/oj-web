package submit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/stretchr/testify/assert"
)

func TestSubmitList(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9001/apiv1/submit/list?problem_id=0&user_name=&status=0&lang=&current_page=1&per_page=10")
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
	assert.Equal(t, 0, respT.Code, "获取失败！")
}
