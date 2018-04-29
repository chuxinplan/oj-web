package apiv1

import (
	"github.com/gin-gonic/gin"
)

func RegisterUser(router *gin.RouterGroup) {
	router.POST("uploadimage", httpHandlerUploadImage)
}

type UserImageParam struct {
	PicType string `form:"pic_type" json:"pic_type"`
}

func httpHandlerUploadImage(c *gin.Context) {

}
