package apiv1

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("login", httpHandlerLogin)
	router.POST("register", httpHandlerRegister)
}

type AccountLogin struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type AccountRegister struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	UserName string `form:"user_name" json:"user_name"`
}

func httpHandlerLogin(c *gin.Context) {
	account := AccountLogin{}
	err := c.Bind(&account)
	if err != nil {
		panic(err)
	}
	id, token, err := managers.AccountLogin(account.Email, account.Password)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    base64.StdEncoding.EncodeToString([]byte(token)),
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, base.Success(id))
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
