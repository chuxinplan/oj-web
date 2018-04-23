package problem

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

type CodeGetParam struct {
	ProblemId int64 `form:"problem_id" json:"problem_id"`
}

func RegisterCode(router *gin.RouterGroup) {
	router.GET("code/get", httpHandlerCodeGet)
	router.POST("code/set", httpHandlerCodeSet)
}

func httpHandlerCodeGet(c *gin.Context) {
	fmt.Println("sssssssssssssssssss")
	userCode := CodeGetParam{}
	err := c.Bind(&userCode)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	code, err := problem.CodeGet(userId, userCode.ProblemId)
	fmt.Println("aaaaaaa", code)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	fmt.Println("bbbbbbbb")
	c.JSON(http.StatusOK, base.Success(code))
}
func httpHandlerCodeSet(c *gin.Context) {
}
