package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/open-fightcoder/oj-web/router/controllers/apiv1/problem"
	"github.com/open-fightcoder/oj-web/router/controllers/apiv1/submit"
)

func Register(router *gin.RouterGroup) {
	RegisterSelf(router)
	RegisterAccount(router)
	RegisterRank(router)

	rankRouter := router.Group("/rank")
	RegisterRank(rankRouter)

	problemRouter := router.Group("/problem")
	problem.RegisterProblem(problemRouter)

	submitRouter := router.Group("/submit")
	submit.RegisterSubmit(submitRouter)

	userRouter := router.Group("/user")
	RegisterUser(userRouter)
}
