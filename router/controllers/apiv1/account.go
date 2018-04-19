package apiv1

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/models"
	"github.com/open-fightcoder/oj-web/router/controllers/baseController"
)

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("account/login", httpHandlerLogin)
	router.POST("account/register", httpHandlerRegister)
}

func httpHandlerLogin(c *gin.Context) {
	account := models.Account{}
	err := c.Bind(&account)
	if err != nil {
		c.JSON(http.StatusOK, (&baseController.Base{}).Fail("参数不合法:", err.Error()))
	}
	if flag, token, mess := managers.AccountLogin(account.Email, account.Password); flag == false {
		c.JSON(http.StatusOK, (&baseController.Base{}).Fail(mess))
	} else {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    base64.StdEncoding.EncodeToString([]byte(token)),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, (&baseController.Base{}).Success())
	}
}

func httpHandlerRegister(c *gin.Context) {
	account := models.Account{}
	err := c.Bind(&account)
	if err != nil {
		c.JSON(http.StatusOK, (&baseController.Base{}).Fail("参数不合法:", err.Error()))
	}
	if flag, userId, mess := managers.AccountRegister(account.Email, account.Password); flag == false {
		c.JSON(http.StatusOK, (&baseController.Base{}).Fail(mess))
	} else {
		c.JSON(http.StatusOK, (&baseController.Base{}).Success(userId))
	}
}
