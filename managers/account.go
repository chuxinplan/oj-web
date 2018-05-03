package managers

import (
	"fmt"

	"io/ioutil"
	"net/http"
	"strings"

	"strconv"
	"time"

	"github.com/open-fightcoder/oj-web/common/components"
	"github.com/open-fightcoder/oj-web/data"
	"github.com/open-fightcoder/oj-web/models"
)

const (
	EMAIL_NOT_EXIT    = 0
	PASSWORD_IS_WRONG = 1
	PARAM_IS_WRONG    = 2
	FIRST_LOGIN       = 3
	LOGIN             = 4
)

func getGithubOpenId(code string) string {
	if code == "" {
		return "-1"
	} else {
		params := "client_id=080191e49e855122ea33&client_secret=34b9a36397b171f01e83fc3c5b676177b29df79e&code="
		params += code
		resp, err := http.Post("https://github.com/login/oauth/access_token",
			"application/x-www-form-urlencoded",
			strings.NewReader(params))
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		strs := strings.Split(string(body), "&")
		token := strings.Split(strs[0], "=")

		url := "https://api.github.com/user?access_token="
		resp, err = http.Get(url + token[1])
		if err != nil {
			panic(err.Error())
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}

		strs = strings.Split(string(body), "\"")
		return strs[3]
	}
}

func getQQOpenId(code string) string {
	if code == "" {
		return "-1"
	} else {
		url := "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=101466300&client_secret=0104260a8f8faac3900cbf184bae55f5&redirect_uri=http%3a%2f%2fxupt4.fightcoder.com%2f%23%2fuser%2flogin&code="
		url += code
		resp, err := http.Get(url)
		if err != nil {
			panic(err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		strs := strings.Split(string(body), "&")
		token := strings.Split(strs[0], "=")

		url = "https://graph.qq.com/oauth2.0/me?access_token="
		resp, err = http.Get(url + token[1])
		if err != nil {
			panic(err.Error())
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}

		strs = strings.Split(string(body), "\"")

		return strs[7]

		//url = "https://graph.qq.com/user/get_user_info?oauth_consumer_key=101466300&access_token=" + token[1] + "&openid=" + strs[7]
		//resp, err = http.Get(url)
		//if err != nil {
		//	panic(err.Error())
		//}
		//
		//defer resp.Body.Close()
		//body, err = ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	panic(err.Error())
		//}
		//mess := &Mess{}
		//if err = json.Unmarshal(body, mess); err != nil {
		//	fmt.Println(err.Error())
		//}
		//c.JSON(http.StatusOK, this.Success(mess))
	}
}

func Login(param1, param2, loginType string) (int, string, int64, string) {
	var accountId int64
	isFirstLogin := false

	if loginType == "simple" {
		account, err := models.AccountGetByEmail(param1)
		if err != nil {
			panic(err)
		}

		if account == nil {
			return EMAIL_NOT_EXIT, "", 0, ""
		} else {
			passwd := account.Password
			if passwd != components.MD5Encode(param2) {
				return PASSWORD_IS_WRONG, "", 0, ""
			}
		}

		accountId = account.Id
	} else if loginType == "qq" {
		openId := getQQOpenId(param1)
		acc, _ := models.AccountGetQQOpenId(openId)
		account := &models.Account{QqId: openId}
		if acc == nil {
			id, _ := models.AccountAdd(account)

			user := &models.User{AccountId: id, NickName: strconv.FormatInt(time.Now().UnixNano(), 10)}
			models.Create(user)
			accountId = id
			isFirstLogin = true
		} else {
			accountId = acc.Id
		}

	} else if loginType == "github" {
		openId := getGithubOpenId(param1)
		acc, _ := models.AccountGetGithubOpenId(openId)
		account := &models.Account{GithubId: openId}
		if acc == nil {
			id, _ := models.AccountAdd(account)

			user := &models.User{AccountId: id, NickName: strconv.FormatInt(time.Now().UnixNano(), 10)}
			models.Create(user)
			accountId = id
			isFirstLogin = true
		} else {
			accountId = acc.Id
		}
	} else {
		return PARAM_IS_WRONG, "", 0, ""
	}

	user, err := models.GetByAccountId(accountId)
	if err != nil {
		fmt.Println(err.Error())
	}
	token, _ := components.CreateToken(user.Id)
	if isFirstLogin {
		return FIRST_LOGIN, token, user.Id, user.UserName
	} else {
		return LOGIN, token, user.Id, user.UserName
	}
}

func AccountRegister(userName string, email string, password string) (int64, error) {
	//TODO 邮箱参数校验,userName校验
	return data.UserRegister(userName, email, components.MD5Encode(password))
}
