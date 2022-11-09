package routes

import (
	controller "bootcamp_es/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

var auth controller.Auth

//	Contains all the routes that are used auth authorize the user or admin

func Authroutes(incommingRoutes *gin.Engine) {
	fmt.Println("authroutess")
	incommingRoutes.GET("/home",auth.Home)
	//	To check wheather the user exist or not
	incommingRoutes.GET("/verifyuser/:username", auth.CheckUser)
	incommingRoutes.GET("/verifyteam/:teamname", auth.CheckTeam)

	//	To send otp to the given number 
	incommingRoutes.POST("/phone_signup", auth.SendPhoneOTP)
	//	To register a new user after checking the otp
	incommingRoutes.POST("/signup", auth.SignUp)
	//	To login an existing user
	incommingRoutes.POST("/login", auth.Login)
	// 	To send otp for changing otp password
	incommingRoutes.POST("/forgot_password", auth.ForgotPassword)
	//	To change the new password
	incommingRoutes.POST("/change_password", auth.ChangePassword)
}

//	Contains all the routes that are used auth search an entity from the landing page

func Search(incommingRoutes *gin.Engine) {
	// To list the user,teams,tournaments,scrims (if any)
	// incommingRoutes.GET("/list_entity", controller.ListEntity)

	//To get the details of the details of the entity
	// incommingRoutes.GET("/seach_entity", controller.GetEntityDetails)

}

