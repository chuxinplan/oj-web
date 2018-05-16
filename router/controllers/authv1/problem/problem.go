package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

type ProblemStatusParam struct {
	ProblemIds string `form:"problem_ids" json:"problem_ids"`
}

func RegisterProblem(router *gin.RouterGroup) {
	router.POST("modify", httpHandlerProblemMod)
	router.POST("delete", httpHandlerProblemDel)
	router.POST("add", httpHandlerProblemAdd)
	router.GET("status", httpHandlerProblemStatus)
}

func httpHandlerProblemStatus(c *gin.Context) {
	param := ProblemStatusParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	mess, err := problem.ProblemStatus(userId, param.ProblemIds)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
func httpHandlerProblemMod(c *gin.Context) {
}

func httpHandlerProblemDel(c *gin.Context) {
}

func httpHandlerProblemAdd(c *gin.Context) {
}
