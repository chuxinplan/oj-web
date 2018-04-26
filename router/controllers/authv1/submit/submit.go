package problem

import "github.com/gin-gonic/gin"

func RegisterSubmit(router *gin.RouterGroup) {
	router.POST("common", httpHandlerSubmitCommon)
	router.POST("test", httpHandlerSubmitTest)
	router.POST("getcommon", httpHandlerSubmitGetCommon)
	router.POST("gettest", httpHandlerSubmitGetTest)
}

func httpHandlerSubmitCommon(c *gin.Context) {
}

func httpHandlerSubmitTest(c *gin.Context) {
}

func httpHandlerSubmitGetCommon(c *gin.Context) {
}

func httpHandlerSubmitGetTest(c *gin.Context) {
}
