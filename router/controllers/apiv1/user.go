package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterUser(router *gin.RouterGroup) {
	router.GET("progress", httpHandlerUserProgress)
	router.GET("getmess", httpHandlerUserGetMess)
	router.GET("count", httpHandlerUserCount)
	router.GET("recentsubmit", httpHandlerUserRecentSubmit)
	router.GET("recentrank", httpHandlerUserRecentRank)
}

type UserParam struct {
	UserName string `form:"user_name" json:"user_name"`
}

func httpHandlerUserRecentSubmit(c *gin.Context) {
	param := UserParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.GetUserRecentSubmit(param.UserName)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerUserRecentRank(c *gin.Context) {
	param := UserParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.GetUserRecentRank(param.UserName)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerUserCount(c *gin.Context) {
	param := UserParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.GetUserCount(param.UserName)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerUserGetMess(c *gin.Context) {
	param := UserParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.GetUserMess(param.UserName)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerUserProgress(c *gin.Context) {
	param := UserParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.GetUserProgress(param.UserName)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
