package login

import (
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/pkg/errors"
)

type Url struct {
}

/**
 * combineURL
 * 拼接url
 * @param string baseURL   基于的url
 * @param map  keysArr     参数列表数组
 * @return string          返回拼接的url
 */
func (this *Url) combineURL(baseURL string, keysArr map[string]string) string {
	str := baseURL + "?"
	for i, v := range keysArr {
		str += i + "=" + v + "&"
	}
	return str[0 : len(str)-1]
}

/**
 * get_contents
 * 服务器通过get请求获得内容
 * @param string url       请求的url,拼接后的
 * @return string          请求返回的内容
 */
func (this *Url) get_contents(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	}
	return "", errors.New("请求错误")
}

/**
 * get
 * get方式请求资源
 * @param string $url     基于的baseUrl
 * @param array $keysArr  参数列表数组
 * @return string         返回的资源内容
 */
func (this *Url) get(url string, keysArr map[string]string) (string, error) {
	requestUrl := this.combineURL(url, keysArr)
	return this.get_contents(requestUrl)
}

/**
 * post_contents
 * 服务器通过post请求获得内容
 * @param string url       请求的url,拼接后的
 * @return string          请求返回的内容
 */
func (this *Url) post_contents(url string, params string) (string, error) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	}
	return "", errors.New("请求错误")
}

/**
 * post
 * post方式请求资源
 * @param string $url       基于的baseUrl
 * @param array $keysArr    请求的参数列表
 * @return string           返回的资源内容
 */
func (this *Url) post(url string, keysArr map[string]string) (string, error) {
	str := ""
	for i, v := range keysArr {
		str += i + "=" + v + "&"
	}
	return this.post_contents(url, str[0:len(str)-1])
}
