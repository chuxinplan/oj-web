package team

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/open-fightcoder/oj-web/managers"
	"net/http"
	"fmt"
)


type TeamInfo struct {
	Id	int64				`form:"id" json:"id"`
	Uid int64				`form:"uid" json:"uid"`
	Name string				`form:"name" json:"name"`
	Description string		`form:"description" json:"description"`
	Avator string			`form:"avator" json:"avator"`
}

type TeamId struct {
	Gid int64 				`form:"gid" json:"gid"`
	Uid int64				`form:"uid" json:"uid"`
}


func RegisterTeam(router * gin.RouterGroup)  {
	router.POST("create", httpHandlerTeamCreate)
	router.POST("modify", httpHandlerTeamUpdate)
	router.POST("delete", httpHandlerTeamDelete)
	router.POST("joinlist", httpHandlerJoinList)
	router.POST("createlist", httpHandlerCreateList)
}

func httpHandlerTeamCreate(c *gin.Context){
	parm := TeamInfo{}
	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	createId, err := managers.TeamCreat(parm.Name, parm.Description, parm.Avator, userId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(createId))
}

func httpHandlerTeamUpdate(c *gin.Context)  {
	parm := TeamInfo{}
	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	status, err := managers.TeamUpdate(parm.Id, userId, parm.Name, parm.Avator, parm.Description)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(status))
}

func httpHandlerTeamDelete(c *gin.Context)  {
	parm := TeamInfo{}

	err := c.Bind(&parm)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	fmt.Println(parm)
	err = managers.TeamRemove(parm.Id, userId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success())
}

func httpHandlerJoinList(c *gin.Context)  {


	userId := base.UserId(c)

	infos, err := managers.TeamsGetByJoin(userId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(infos))
}

func httpHandlerCreateList(c *gin.Context) {
	userId := base.UserId(c)

	infos, err := managers.TeamsGetByCreate(userId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(infos))

}