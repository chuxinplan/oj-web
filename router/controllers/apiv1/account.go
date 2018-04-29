package apiv1

import (
	"encoding/base64"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/pkg/errors"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("login", httpHandlerLogin)
	router.POST("register", httpHandlerRegister)
}

type LoginTypeParam struct {
	Type string `form:"type" json:"type"`
}

type AccountSimpleLogin struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type AccountOtherLogin struct {
	Code  string `form:"code" json:"code"`
	State string `form:"state" json:"state"`
}

type AccountRegister struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	UserName string `form:"user_name" json:"user_name"`
}

func httpHandlerLogin(c *gin.Context) {
	param := LoginTypeParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	loginType := param.Type
	var state int
	var msg string
	var userId int64
	if loginType == "qq" || loginType == "github" {
		account := AccountOtherLogin{}
		err := c.Bind(&account)
		if err != nil {
			panic(err)
		}
		state, msg, userId = managers.Login(account.Code, account.State, loginType)
	} else if loginType == "simple" {
		account := AccountSimpleLogin{}
		err := c.Bind(&account)
		if err != nil {
			panic(err)
		}
		state, msg, userId = managers.Login(account.Email, account.Password, loginType)
	} else {
		panic(errors.New("参数错误"))
	}

	if state == managers.EMAIL_NOT_EXIT || state == managers.PASSWORD_IS_WRONG || state == managers.PARAM_IS_WRONG {
		var msg string
		switch state {
		case managers.EMAIL_NOT_EXIT:
			msg = "Email not exit!"
			break
		case managers.PASSWORD_IS_WRONG:
			msg = "Password is wrong!"
			break
		case managers.PARAM_IS_WRONG:
			msg = "Param is wrong!"
			break
		}
		c.JSON(http.StatusOK, base.Fail(msg))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(msg)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		result := make(map[string]string)
		if state == managers.FIRST_LOGIN {
			result["is_first"] = "true"
			result["user_id"] = strconv.FormatInt(userId, 10)
		} else {
			result["is_first"] = "false"
			result["user_id"] = strconv.FormatInt(userId, 10)
		}
		c.JSON(http.StatusOK, base.Success(result))
	}
}

func httpHandlerRegister(c *gin.Context) {
	account := AccountRegister{}
	err := c.Bind(&account)
	if err != nil {
		panic(err)
	}
	userId, err := managers.AccountRegister(account.UserName, account.Email, account.Password)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(userId))
}
