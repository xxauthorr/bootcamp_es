package routes

import (
	"bootcamp_es/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	team controllers.Team
)

func Team(incommingRoutes *gin.Engine) {
	fmt.Println("hereeeee")
	incommingRoutes.Use(mw.Authneticate)
	incommingRoutes.POST("/team_reg", team.RegisterTeam)

}
