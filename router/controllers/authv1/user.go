package authv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterUser(router *gin.RouterGroup) {
	router.POST("uploadimage", httpHandlerUploadImage)
}

type UserImageParam struct {
	PicType string `form:"pic_type" json:"pic_type"`
}

func httpHandlerUploadImage(c *gin.Context) {
	param := UserImageParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	file, _, _ := c.Request.FormFile("upload")
	userId := base.UserId(c)
	err = managers.UploadImage(file, userId, param.PicType)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success("上传成功"))
}
