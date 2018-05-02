package team

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/open-fightcoder/oj-web/managers"
	"net/http"
	"fmt"
)

type Member struct {
	Gid int64	`form:"gid" json:"gid"`
	Uid int64	`form:"uid" json:"uid"`
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

	fmt.Println(parm)

	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}

	fmt.Println(parm)

	Userid := base.UserId(c)

	status, err := managers.MemberAdd(parm.Gid, parm.Uid, Userid)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(status))
}


