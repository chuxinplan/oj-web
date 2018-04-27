package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

type CodeGetParam struct {
	ProblemId int64 `form:"problem_id" json:"problem_id"`
}

type CodeSetParam struct {
	ProblemId int64  `form:"problem_id" json:"problem_id"`
	SaveCode  string `form:"save_code" json:"save_code"`
	Language  string `form:"language" json:"language"`
}

func RegisterCode(router *gin.RouterGroup) {
	router.GET("code/get", httpHandlerCodeGet)
	router.POST("code/set", httpHandlerCodeSet)
}

func httpHandlerCodeGet(c *gin.Context) {
	userCode := CodeGetParam{}
	err := c.Bind(&userCode)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	code, err := problem.CodeGet(userId, userCode.ProblemId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(code))
}
func httpHandlerCodeSet(c *gin.Context) {
	userCode := CodeSetParam{}
	err := c.Bind(&userCode)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	err = problem.CodeSet(userCode.ProblemId, userId, userCode.SaveCode, userCode.Language)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success())
}
