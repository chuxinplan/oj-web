package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterUser(router *gin.RouterGroup) {
	router.GET("progress", httpHandlerUserProgress)
}

type UserImageParam struct {
	UserId int64 `form:"user_id" json:"user_id"`
}

func httpHandlerUserProgress(c *gin.Context) {
	param := UserImageParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	mess, err := managers.GetUserProgress(param.UserId)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(mess))
}
