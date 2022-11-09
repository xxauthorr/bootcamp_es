package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

var mw middlewares.Mwares
var edit controllers.UserEdit
var user controllers.User

// Contains all the routes that are used for the user profile (not to edit)

func User(incommingRoutes *gin.Engine) {
	route := incommingRoutes.Group("/user")
	route.Use(mw.AuthneticateToken)
	route.GET("/:username", user.UserProfile)
	route.PUT("/editbio", edit.BioEdit)
	route.PUT("/editsocial", edit.UserSocialEdit)
	route.PUT("/addachievements", edit.UserAcheivementsAdd)
	route.DELETE("/delachievements", edit.UserAcheivementsDelete)
	route.PUT("/updatenotification",edit.UpdateNotification)

}

// Contains all the routes to edit the control settings (password,email,phone)
// func UserSettings(incommingRoutes *gin.Engine) {

// }
