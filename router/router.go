package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/common/g"
	apiv1 "github.com/open-fightcoder/oj-web/router/controllers/api/v1"
	authv1 "github.com/open-fightcoder/oj-web/router/controllers/auth/v1"
	"github.com/open-fightcoder/oj-web/router/middleware"
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
	router = gin.Default()

	router.Use(middleware.Cors())
	router.Use(middleware.Recovery())
	router.Use(middleware.MaxAllowed(g.Conf().Run.MaxAllowed))

	authRouter := router.Group("/auth", middleware.Auth())
	authv1.RegisterAUTHV1(authRouter)

	apiRouter := router.Group("/api")
	apiv1.RegisterAPIV1(apiRouter)
}
