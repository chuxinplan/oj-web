package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/apiv1/problem"
)

func Register(router *gin.RouterGroup) {
	RegisterSelf(router)
	RegisterAccount(router)
	RegisterRank(router)

	problemRouter := router.Group("/problem")
	problem.RegisterProblem(problemRouter)

}
