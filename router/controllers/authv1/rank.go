package authv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterRank(router *gin.RouterGroup) {
	router.GET("getlist", httpHandlerGetList)
}

type RankParam struct {
	IsWeek int `form:"is_week" json:"is_week"`
}

func httpHandlerGetList(c *gin.Context) {
	param := RankParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.RankGetList(param.IsWeek)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
