package main

import (
	"net/http"

	"time"

	"github.com/appleboy/gin-revision-middleware"
	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	_ "github.com/open-fightcoder/oj-web/common"
	"github.com/sirupsen/logrus"
)

func rootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Hello": "world",
	})
	// logrus.SetFormatter()
	context.JSON(200, gin.H{"token": "aaa"})
}

func main() {

	router := gin.New()
	gin.Logger()
	router.Use(cors.Default())
	//修订中间件：显示版本信息在RespHeader
	router.Use(revision.Middleware())
	//限制最大请求并发量
	router.Use(limit.MaxAllowed(20))
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	router.GET("/", rootHandler)
	router.Run(":" + "8000")
}
