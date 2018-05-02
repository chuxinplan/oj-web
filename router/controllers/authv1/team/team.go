package team

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"net/http"
	"github.com/open-fightcoder/oj-web/managers"
)


type TeamInfo struct {
	Id	int64				`form:"id" json:"id"`
	Uid int64				`form:"uid" json:"uid"`
	Name string				`form:"name" json:"name"`
	Description string		`form:"description" json:"description"`
	Avator string			`form:"avator" json:"avator"`
}



func RegisterTeam(router * gin.RouterGroup)  {
	router.POST("create", httpHandlerTeamCreate)
	router.POST("modify", httpHandlerTeamUpdate)
	router.POST("delete", httpHandlerTeamDelete)
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
	err = managers.TeamRemove(parm.Id, userId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success())
}
