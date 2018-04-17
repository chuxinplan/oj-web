package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/common/g"
)

func RegisterSelf(router *gin.RouterGroup) {
	router.GET("/self/health", httpHandlerHealth)
	router.GET("/self/config", httpHandlerConfig)
}

func httpHandlerHealth(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func httpHandlerConfig(c *gin.Context) {
	c.String(http.StatusOK, g.Conf())
}
