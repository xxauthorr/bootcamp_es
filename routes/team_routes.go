package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type team struct {
	cTeamEdit   controllers.EditTeam
	teamMwares middlewares.TeamCheckers
	mw          middlewares.Mwares
}

var te team

func Team(incommingRoutes *gin.Engine) {
	routes := incommingRoutes.Group("/team")
	routes.Use(te.mw.AuthneticateToken, te.teamMwares.CheckLeaderTeam)
	routes.PUT("/editbio", te.cTeamEdit.TeamEditBio)
	routes.PUT("/addachievements/:content", te.cTeamEdit.TeamAddAchievements)
	routes.DELETE("/delachievements", te.cTeamEdit.TeamDelAchievements)
}