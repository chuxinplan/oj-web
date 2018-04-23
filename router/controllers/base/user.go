package base

import "github.com/gin-gonic/gin"

func UserId(c *gin.Context) int64 {
	id, _ := c.Get("userId")
	return id.(int64)
}
