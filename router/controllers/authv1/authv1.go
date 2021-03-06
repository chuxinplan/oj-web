package authv1

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/authv1/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/authv1/submit"
	"github.com/open-fightcoder/oj-web/router/controllers/authv1/team"
)

func Register(router *gin.RouterGroup) {
	problemRouter := router.Group("/problem")
	problem.RegisterProblem(problemRouter)
	problem.RegisterCode(problemRouter)
	problem.RegisterCollection(problemRouter)
	problem.RegisterUserProblem(problemRouter)

	teamRouter := router.Group("/team")
	team.RegisterTeam(teamRouter)
	team.RegisterMember(teamRouter)

	RegisterAccount(router)
	rankRouter := router.Group("/rank")
	RegisterRank(rankRouter)

	submitRouter := router.Group("/submit")
	submit.RegisterSubmit(submitRouter)

	userRouter := router.Group("/user")
	RegisterUser(userRouter)
}
