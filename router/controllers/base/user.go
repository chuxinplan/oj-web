package base

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserId(c *gin.Context) int64 {
	str, _ := c.Get("userId")
	id, _ := strconv.ParseInt(str.(string), 10, 64)
	return id
}
