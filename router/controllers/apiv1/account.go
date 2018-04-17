package apiv1

import "github.com/gin-gonic/gin"

func RegisterAccount(router *gin.RouterGroup) {
	router.POST("account/login", httpHandlerLogin)
	router.POST("account/register", httpHandlerRegister)
}

func httpHandlerLogin(c *gin.Context) {
}

func httpHandlerRegister(c *gin.Context) {

}
