package authv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/router/controllers/base"
)

func RegisterUser(router *gin.RouterGroup) {
	router.POST("uploadimage", httpHandlerUploadImage)
	router.POST("updatemess", httpHandlerUpdateMess)
	router.POST("collection", httpHandlerCollection)
}

type UserImageParam struct {
	PicType string `form:"pic_type" json:"pic_type"`
}

type UserMessParam struct {
	UserName     string `form:"user_name" json:"user_name"`
	NickName     string `form:"nick_name" json:"nick_name"`
	Sex          string `form:"sex" json:"sex"`
	Blog         string `form:"blog" json:"blog"`
	Git          string `form:"git" json:"git"`
	Description  string `form:"description" json:"description"`
	Birthday     string `form:"birthday" json:"birthday"`
	DailyAddress string `form:"daily_address" json:"daily_address"`
	StatSchool   string `form:"stat_school" json:"stat_school"`
	SchoolName   string `form:"school_name" json:"school_name"`
}

func httpHandlerCollection(c *gin.Context) {

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

func httpHandlerUpdateMess(c *gin.Context) {
	param := UserMessParam{}
	err := c.Bind(&param)
	if err != nil {
		panic(err)
	}
	userId := base.UserId(c)
	err = managers.UpdateUserMess(userId, param.UserName, param.NickName, param.Sex, param.Blog, param.Git, param.Description, param.Birthday, param.DailyAddress, param.StatSchool, param.SchoolName)
	if err != nil {
		c.JSON(http.StatusOK, base.Fail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, base.Success(""))
}
