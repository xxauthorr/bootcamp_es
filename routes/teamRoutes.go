package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

var ()

type team struct {
	cTeamEdit   controllers.EditTeam
	teamChecker middlewares.TeamCheckers
	mw          middlewares.Mwares
}

var te team

func Team(incommingRoutes *gin.Engine) {
	routes := incommingRoutes.Group("/team")
	routes.Use(te.mw.AuthneticateToken, te.teamChecker.CheckLeaderTeam)
	routes.PUT("/editbio", te.cTeamEdit.TeamEditBio)
	routes.PUT("/addachievements", te.cTeamEdit.TeamAddAchievements)
	routes.DELETE("/delachievements", te.cTeamEdit.TeamDelAchievements)
}
