package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type tournament struct {
	controller  controllers.Tour
	mw          middlewares.Mwares
	tcheck      middlewares.TeamCheckers
	tourChecker middlewares.TournamentChecker
}

var t tournament

func Tournament(incommingRoutes *gin.Engine) {
	routes := incommingRoutes.Group("/tournament")
	tourOperations := incommingRoutes.Group("/tournament/edit")
	// routes.Use(t.mw.AuthneticateToken, t.tcheck.CheckLeaderTeam, t.tcheck.CheckTeamStrength)
	routes.Use(t.mw.AuthneticateToken, t.tcheck.CheckLeaderTeam)
	routes.POST("/registration", t.controller.TournamentRegistration)
	tourOperations.Use(t.mw.AuthneticateToken, t.tourChecker.CheckOwner)
	tourOperations.PUT("/edittournament", t.controller.EditTournamentData)
}
