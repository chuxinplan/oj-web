package team

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/open-fightcoder/oj-web/managers"
	"net/http"
)

type Member struct {
	Gid int64	`form:"gid" json:"gid"`
	Uid int64	`form:"uid" json:"uid"`
	Gname string `form:"gname" json:"gname"`
	Uname string `form:"uname" json:"uname"`
}

func RegisterMember(router *gin.RouterGroup)  {
	router.POST("member/add", httpHandlerMemeberAdd)
	router.POST("member/delete", httpHandlerMemberDel)
}

func httpHandlerMemberDel(c *gin.Context)  {
	parm := Member{}

	if err := c.Bind(&parm); err != nil {
		panic(err)
	}

	Userid := base.UserId(c)
	err := managers.MemberRemove(parm.Uid, parm.Gid, Userid)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success())
}


func httpHandlerMemeberAdd(c *gin.Context)  {
	parm := Member{}

	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}

	Userid := base.UserId(c)

	status, err := managers.MemberAdd(parm.Gid, parm.Uid, Userid)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(status))
}


