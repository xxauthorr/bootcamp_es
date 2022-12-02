package routes

import (
	"net/http"

	controller "github.com/xxauthorr/bootcamp_es/controllers"

	"github.com/gin-gonic/gin"
)

type unAuth struct {
	c    controller.Auth
	team controller.Team
	user controller.User
}

var auth unAuth

//	Contains all the routes that are used auth authorize the user or admin

func AuthRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/", auth.c.Home)
	incommingRoutes.GET("/:username", auth.user.UserProfile)
	incommingRoutes.GET("/teamprofile/:teamname", auth.team.TeamProfile)
	incommingRoutes.GET("/searchfive/:entity", auth.c.SearchFirstFive)
	incommingRoutes.GET("/tournamentprofile/:tournament", auth.c.ShowTournament)
	incommingRoutes.GET("/verifyuser/:username", auth.c.CheckUser)
	incommingRoutes.GET("/verifyteam/:teamname", auth.c.CheckTeam)

	
	var routes = incommingRoutes.Group("/auth")
	routes.GET("/gettoken", auth.c.GetToken)
	routes.POST("/login", auth.c.Login)
	routes.POST("/otp", auth.c.SendPhoneOTP)
	routes.POST("/signup", auth.c.SignUp)
	routes.POST("/forgotpassword", auth.c.ForgotPassword)
	routes.POST("/changepassword", auth.c.ChangePassword)

}


func BadRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"status": false, "message": "page not found"})
}
