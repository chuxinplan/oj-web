package problem

import "github.com/gin-gonic/gin"

func RegisterProblem(router *gin.RouterGroup) {
	router.POST("modify", httpHandlerProblemMod)
	router.POST("delete", httpHandlerProblemDel)
	router.POST("add", httpHandlerProblemAdd)
}

func httpHandlerProblemMod(c *gin.Context) {
}

func httpHandlerProblemDel(c *gin.Context) {
}

func httpHandlerProblemAdd(c *gin.Context) {
}
