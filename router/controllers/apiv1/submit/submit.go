package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterSubmit(router *gin.RouterGroup) {
	router.GET("list", httpHandlerSubmitList)
}

type ListParam struct {
	ProblemId   int64  `form:"problem_id" json:"problem_id"`
	UserName    string `form:"user_name" json:"user_name"`
	Status      int    `form:"status" json:"status"`
	Lang        string `form:"lang" json:"lang"`
	CurrentPage int    `form:"current_page" json:"current_page"`
	PerPage     int    `form:"per_page" json:"per_page"`
}

func httpHandlerSubmitList(c *gin.Context) {
	param := ListParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := problem.SubmitList(param.Origin, param.Tag, param.Sort, param.IsAsc, param.CurrentPage, param.PerPage)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
