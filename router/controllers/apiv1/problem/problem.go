package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

type ListParam struct {
	Origin      string `form:"origin" json:"origin"`
	Tag         string `form:"tag" json:"tag"`
	Difficult   string `form:"diff" json:"diff"`
	Sort        int    `form:"sort" json:"sort"`
	IsAsc       int    `form:"is_asc" json:"is_asc"`
	CurrentPage int    `form:"current_page" json:"current_page"`
	PerPage     int    `form:"per_page" json:"per_page"`
}

type GetParam struct {
	Id int64 `form:"id" json:"id"`
}

type RandomParam struct {
	Origin string `form:"origin" json:"origin"`
	Tag    string `form:"tag" json:"tag"`
}

func RegisterProblem(router *gin.RouterGroup) {
	router.GET("list", httpHandlerProblemList)
	router.GET("get", httpHandlerProblemGet)
	router.GET("random", httpHandlerProblemRandom)
}

func httpHandlerProblemList(c *gin.Context) {
	param := ListParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := problem.ProblemList(param.Difficult, param.Origin, param.Tag, param.Sort, param.IsAsc, param.CurrentPage, param.PerPage)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerProblemGet(c *gin.Context) {
	param := GetParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	problemMess, err := problem.ProblemGet(param.Id)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(problemMess))
}

func httpHandlerProblemRandom(c *gin.Context) {
	param := RandomParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	problemMess, err := problem.ProblemRandom(param.Origin, param.Tag)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(problemMess))
}
