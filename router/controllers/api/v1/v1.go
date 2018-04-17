package controllers

import "github.com/gin-gonic/gin"

func RegisterAPIV1(router *gin.RouterGroup) {
	RegisterSelf(router)

}
