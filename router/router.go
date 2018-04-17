package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	apiv1 "github.com/open-fightcoder/oj-web/router/controllers/api/v1"
	authv1 "github.com/open-fightcoder/oj-web/router/controllers/auth/v1"
)

var router *gin.Engine
var once sync.Once

// 获取路由并初始化
func GetRouter() *gin.Engine {
	// 只执行一次
	once.Do(func() {
		initRouter()
	})
	return router
}

// 初始化路由
func initRouter() {
	router := gin.New()

	authRouter := router.Group("/auth")
	authv1.RegisterAUTHV1(authRouter)

	apiRouter := router.Group("/api")
	apiv1.RegisterAPIV1(apiRouter)
}
