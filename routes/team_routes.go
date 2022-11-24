package routes

import (
	"github.com/xxauthorr/bootcamp_es/controllers"
	"github.com/xxauthorr/bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type team struct {
	cTeamEdit  controllers.EditTeam
	teamMwares middlewares.TeamCheckers
	mw         middlewares.Mwares
}

var te team

func Team(incommingRoutes *gin.Engine) {
	routes := incommingRoutes.Group("/team")
	routes.Use(te.mw.AuthneticateToken, te.teamMwares.CheckLeaderTeam)
	routes.PUT("/editbio", te.cTeamEdit.TeamEditBio)
	routes.PUT("/addachievements/:content", te.cTeamEdit.TeamAddAchievements)
	routes.DELETE("/delachievements", te.cTeamEdit.TeamDelAchievements)
	routes.PUT("/updatenotification/:action", te.cTeamEdit.UpdateTeamNotification)
	routes.PUT("/joinrequest/:user",te.cTeamEdit.SendTeamJoinRequeset)
}
