package routes

import (
	"bootcamp_es/controllers"

	"github.com/gin-gonic/gin"
)

var (
	team controllers.Team
)

func Team(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(mw.Authneticate)
	incommingRoutes.POST("/team_reg", team.RegisterTeam)

}
