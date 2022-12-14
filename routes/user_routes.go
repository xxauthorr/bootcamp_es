package routes

import (
	"github.com/xxauthorr/bootcamp_es/controllers"
	"github.com/xxauthorr/bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type user struct {
	cTeam controllers.Team
	mw    middlewares.Mwares
	edit  controllers.UserEdit
}

var us user

// Contains all the routes that are used for the user profile (not to edit)

func User(incommingRoutes *gin.Engine) {
	route := incommingRoutes.Group("/user")
	route.Use(us.mw.AuthneticateToken)
	route.PUT("/editbio", us.edit.BioEdit)
	route.PUT("/editsocial", us.edit.UserSocialEdit)
	route.PUT("/addachievements", us.edit.UserAddAcheivements)
	route.DELETE("/delachievements", us.edit.UserDelAcheivements)
	route.PUT("/updatenotification/:action", us.edit.UpdateNotification)
	route.POST("/teamregistration", us.cTeam.RegisterTeam)
	route.PUT("/userpops/:to", us.edit.UserPopularityEdit)
	route.PUT("/teamrequest/:team", us.edit.SendTeamJoinRequest)
	route.PUT("/teamexit", us.edit.ExitTeam)
}

// Contains all the routes to edit the control settings (password,email,phone)
// func UserSettings(incommingRoutes *gin.Engine) {

// }
