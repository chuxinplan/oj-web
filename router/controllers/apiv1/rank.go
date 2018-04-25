package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterRank(router *gin.RouterGroup) {
	router.GET("get", httpHandlerGet)
}

type RankParam struct {
	CurrentPage int `form:"current_page" json:"current_page"`
	PerPage     int `form:"per_page" json:"per_page"`
}

func httpHandlerGet(c *gin.Context) {
	param := RankParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.RankGet(param.CurrentPage, param.PerPage)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
