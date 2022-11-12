package routes

import (
	controller "bootcamp_es/controllers"

	"github.com/gin-gonic/gin"
)

type unAuth struct {
	c    controller.Auth
	user controller.User
}

var auth unAuth

//	Contains all the routes that are used auth authorize the user or admin

func Authroutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/", auth.c.Home)
	incommingRoutes.GET("/:username", auth.user.UserProfile)

	//	To check wheather the user exist or not
	incommingRoutes.GET("/verifyuser/:username", auth.c.CheckUser)
	incommingRoutes.GET("/verifyteam/:teamname", auth.c.CheckTeam)
	routes := incommingRoutes.Group("/auth")
	routes.POST("/login", auth.c.Login)
	routes.POST("/otp", auth.c.SendPhoneOTP)
	routes.POST("/signup", auth.c.SignUp)
	routes.POST("/forgotpassword", auth.c.ForgotPassword)
	routes.POST("/changepassword", auth.c.ChangePassword)
}

func Profiles(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/userprofile/:username", auth.c.UserProfile)
	incommingRoutes.GET("/teamprofile/:username", auth.c.TeamProfile)
}

//	Contains all the routes that are used auth search an entity from the landing page

func Search(incommingRoutes *gin.Engine) {
	// To list the user,teams,tournaments,scrims (if any)
	// incommingRoutes.GET("/list_entity", controller.ListEntity)

	//To get the details of the details of the entity
	// incommingRoutes.GET("/seach_entity", controller.GetEntityDetails)

}
