package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	team     controllers.Team
	checker middlewares.Checkers
)

func Team(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(mw.AuthneticateToken)
	incommingRoutes.POST("/team_reg", team.RegisterTeam)
	routes := incommingRoutes.Group("/team")
	routes.Use(checker.CheckLeaderTeam)
}
