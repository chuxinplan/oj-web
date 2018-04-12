package routers

import "github.com/gin-gonic/gin"

var router *gin.Engine

func GetRouter() *gin.Engine {
	initRouter()
	return router
}

func initRouter() {

}
