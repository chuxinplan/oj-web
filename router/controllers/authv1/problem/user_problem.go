package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

type UserProblemListParam struct {
	CurrentPage int `form:"current_page" json:"current_page"`
	PerPage     int `form:"per_page" json:"per_page"`
}

func RegisterUserProblem(router *gin.RouterGroup) {
	router.GET("userproblem/list", httpHandlerUserProblemList)
	router.GET("userproblem/get", httpHandlerUserProblemGet)
	router.POST("userproblem/check", httpHandlerUserProblemCheck)
	router.POST("userproblem/add", httpHandlerUserProblemAdd)
	router.POST("userproblem/modify", httpHandlerUserProblemMod)

}
func httpHandlerUserProblemList(c *gin.Context) {
	userProblem := UserProblemListParam{}
	err := c.Bind(&userProblem)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	problemList, err := problem.UserProblemList(userId, userProblem.CurrentPage, userProblem.PerPage)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(problemList))
}
func httpHandlerUserProblemGet(c *gin.Context) {

}
func httpHandlerUserProblemCheck(c *gin.Context) {
}

func httpHandlerUserProblemAdd(c *gin.Context) {
}
func httpHandlerUserProblemMod(c *gin.Context) {
}
