package problem

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
	"github.com/pkg/errors"
)

type CollectionGetParam struct {
	ProblemIds string `form:"problem_ids" json:"problem_ids"`
}

type CollectionSetParam struct {
	ProblemId int64  `form:"problem_id" json:"problem_id"`
	Flag      string `form:"flag" json:"flag"`
}

func RegisterCollection(router *gin.RouterGroup) {
	router.GET("collection/get", httpHandlerCollectionGet)
	router.POST("collection/set", httpHandlerCollectionSet)
}

func httpHandlerCollectionGet(c *gin.Context) {
	param := CollectionGetParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	mess, err := problem.CollectionGet(userId, param.ProblemIds)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
func httpHandlerCollectionSet(c *gin.Context) {
	userCollection := CollectionSetParam{}
	err := c.Bind(&userCollection)
	if err != nil {
		panic(err)
	}
	if userCollection.Flag != "set" && userCollection.Flag != "cancel" {
		panic(errors.New("参数不符合"))
	}
	userId := base.UserId(c)
	mess, err := problem.CollectionSet(userId, userCollection.ProblemId, userCollection.Flag)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
