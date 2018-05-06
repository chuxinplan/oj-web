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
	Gname string `form:"gname" json:"gname"`
	Uname string `form:"uname" json:"uname"`
	Stats int `form:"stats" json:"stats"` //申请为true， 添加为flase
}


func RegisterMember(router *gin.RouterGroup)  {
	router.POST("member/add/invite", httpHandlerMemberInvite)
	router.POST("member/add/apply", httpHandlerMemeberApply)
	router.POST("member/add/audit",  httpHandlerMemeberAudit)
	router.POST("member/add/accept",  httpHandlerMemeberAccept)
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


func httpHandlerMemeberAccept(c *gin.Context)  {
	parm := Member{}

	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}

	Userid := base.UserId(c)

	status, err := managers.MemberAccept(parm.Gid , Userid, parm.Stats)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(status))
}


func httpHandlerMemeberApply(c *gin.Context){
	parm := Member{}

	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}

	Userid := base.UserId(c)
	fmt.Println(Userid)

	status, err := managers.MemberApply(parm.Gid , Userid)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(status))

}

func httpHandlerMemeberAudit(c *gin.Context) {
	parm := Member{}

	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}

	Userid := base.UserId(c)

	status, err := managers.MemberAudit(parm.Gid , parm.Uid, Userid, parm.Stats)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(status))

}

func httpHandlerMemberInvite(c *gin.Context) {
	parm := Member{}

	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}

	Userid := base.UserId(c)

	status, err := managers.MemberInvite(parm.Gid , parm.Uid, Userid)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(status))

}