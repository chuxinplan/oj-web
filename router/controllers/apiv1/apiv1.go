package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/apiv1/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/apiv1/team"
)

func Register(router *gin.RouterGroup) {
	RegisterSelf(router)
	RegisterAccount(router)

	problemRouter := router.Group("/problem")
	problem.RegisterProblem(problemRouter)

	teamRouter := router.Group("/team")
	team.RegisterTeam(teamRouter)

}
