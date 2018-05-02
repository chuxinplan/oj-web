package team

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"net/http"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"fmt"
)


type Team struct {
	Gid int64 	`form:"gid" json:"gid"`
}

func RegisterTeam(router *gin.RouterGroup)  {
	router.GET("info", htttpHandlerTeamGet)
	router.GET("member", httpHandlerMember)
}

func htttpHandlerTeamGet(c *gin.Context) {

	param := Team{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}

	teaminfo, err := managers.GetTeam(param.Gid)

	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.Success(teaminfo))

}

func httpHandlerMember(c *gin.Context) {

}
