package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type tournament struct {
	controller controllers.Tour
	mw         middlewares.Mwares
	check      middlewares.TeamCheckers
}

var t tournament

func Tournament(incommingRoutes *gin.Engine) {
	routes := incommingRoutes.Group("tournament")
	routes.Use(t.mw.AuthneticateToken, t.check.CheckLeaderTeam,t.check.CheckTeamStrength)
	// routes.Use(t.mw.AuthneticateToken, t.check.CheckLeaderTeam)
	routes.POST("/registration", t.controller.TournamentRegistration)
}
