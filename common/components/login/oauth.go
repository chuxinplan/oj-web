package login

import (
	"encoding/json"
	"math/rand"
	"net/url"
	"strconv"

	"errors"

	"strings"

	"github.com/open-fightcoder/oj-web/common/components"
)

type ErrorMess struct {
	Error            int64  `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type OpenIdMess struct {
	ClientId string `json:"client_id"`
	OpenId   string `json:"openid"`
}

//生成state，拼接请求URL
func QQLogin() string {
	randNum := rand.Intn(1000)
	state := components.MD5Encode(strconv.Itoa(randNum))
	//TODO 将state添加到Session
	baseUrl := "https://graph.qq.com/oauth2.0/authorize"
	param := map[string]string{
		"response_type": "code",
		"client_id":     "101466300",
		"redirect_uri":  url.QueryEscape("http://www.fightcoder.com/#/user/login"),
		"state":         state,
	}
	return (&Url{}).combineURL(baseUrl, param)
}

//返回access_token
func QQCallback(code string, reqState string) (string, error) {
	//成功：access_token=28AA149D4520BAA0EA7A09879B81A3DE&expires_in=7776000&refresh_token=B9D9DED6BBAC973EDF0FD51B7AF8362F
	//失败：callback( {"error":100020,"error_description":"code is reused error"} );
	//TODO Session取出state
	//state := "aaassda"
	//if reqState != state {
	//	return "", errors.New("The state does not match. You may be a victim of CSRF.")
	//}
	baseUrl := "https://graph.qq.com/oauth2.0/token"
	param := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     "101466300",
		"client_secret": "0104260a8f8faac3900cbf184bae55f5",
		"redirect_uri":  url.QueryEscape("http://www.fightcoder.com/#/user/login"),
		"code":          code,
	}
	body, err := (&Url{}).get(baseUrl, param)
	if err != nil {
		return "", errors.New("get response fail")
	}
	if strings.Contains(body, "callback") {
		start := strings.Index(body, "(")
		end := strings.Index(body, ")")
		body = body[start+1 : end]
		mess := &ErrorMess{}
		err := json.Unmarshal([]byte(body), mess)
		if err != nil {
			return "", errors.New("decode response fail")
		}
		return "", errors.New(mess.ErrorDescription)
	}
	start := strings.Index(body, "=")
	end := strings.Index(body, "&")
	return body[start+1 : end], nil
}

//返回openId
func GetOpenid(accessToken string) (string, error) {
	//成功：callback( {"client_id":"101466300","openid":"9EC58C000E554465E68F8F51D3D1A1AF"} );
	//失败：callback( {"error":100013,"error_description":"access token is illegal"} );
	baseUrl := "https://graph.qq.com/oauth2.0/me"
	param := map[string]string{
		"access_token": accessToken,
	}
	body, err := (&Url{}).get(baseUrl, param)
	if err != nil {
		return "", errors.New("get response fail")
	}
	if strings.Contains(body, "callback") {
		start := strings.Index(body, "(")
		end := strings.Index(body, ")")
		body = body[start+1 : end]
	}
	if strings.Contains(body, "error") {
		mess := &ErrorMess{}
		err = json.Unmarshal([]byte(body), mess)
		if err != nil {
			return "", errors.New("decode response fail")
		}
		return "", errors.New(mess.ErrorDescription)
	} else {
		mess := &OpenIdMess{}
		err = json.Unmarshal([]byte(body), mess)
		if err != nil {
			return "", errors.New("decode response fail")
		}
		return mess.OpenId, nil
	}
}
