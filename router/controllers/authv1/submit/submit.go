package submit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers/submit"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterSubmit(router *gin.RouterGroup) {
	router.POST("common", httpHandlerSubmitCommon)
	router.POST("test", httpHandlerSubmitTest)
	router.GET("getcommon", httpHandlerSubmitGetCommon)
	router.GET("gettest", httpHandlerSubmitGetTest)
}

type CommonParam struct {
	ProblemId int64  `form:"problem_id" json:"problem_id"`
	Language  string `form:"language" json:"language"`
	Code      string `form:"code" json:"code"`
}

type TestParam struct {
	Language string `form:"language" json:"language"`
	Input    string `form:"input" json:"input"`
	Code     string `form:"code" json:"code"`
}

type GetParam struct {
	SubmitId int64 `form:"submit_id" json:"submit_id"`
}

func httpHandlerSubmitCommon(c *gin.Context) {
	param := CommonParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	mess, err := submit.SubmitCommon(param.ProblemId, userId, param.Language, param.Code)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerSubmitTest(c *gin.Context) {
	param := TestParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	mess, err := submit.SubmitTest(userId, param.Language, param.Input, param.Code)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerSubmitGetCommon(c *gin.Context) {
	param := GetParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := submit.SubmitGetCommon(param.SubmitId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}

func httpHandlerSubmitGetTest(c *gin.Context) {
	param := GetParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := submit.SubmitGetTest(param.SubmitId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
