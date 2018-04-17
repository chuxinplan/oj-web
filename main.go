package main

import (
	"flag"

	"github.com/open-fightcoder/oj-web/common"
	"github.com/open-fightcoder/oj-web/router"
)

//func rootHandler(context *gin.Context) {
//	context.JSON(http.StatusOK, gin.H{
//		"Hello": "world",
//	})
//	// logrus.SetFormatter()
//	context.JSON(200, gin.H{"token": "aaa"})
//}

func main() {
	//version := flag.Bool("v", false, "show version")
	cfgFile := flag.String("c", "cfg/cfg.toml.debug", "set config file")
	flag.Parse()

	common.Init(*cfgFile)

	router := router.GetRouter()

	//gin.Logger()
	//router.Use(cors.Default())
	////修订中间件：显示版本信息在RespHeader
	//router.Use(revision.Middleware())
	////限制最大请求并发量
	//router.Use(limit.MaxAllowed(20))
	//router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	router.Run(":" + "8000")
}
