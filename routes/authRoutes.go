package routes

import (
	controller "bootcamp_es/controllers"

	"github.com/gin-gonic/gin"
)

var do controller.Controller

//	Contains all the routes that are used do authorize the user or admin

func Authroutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/send_otp", do.SendPhoneOTP)
	incommingRoutes.GET("/check_otp", do.CheckPhoneOtp)
	incommingRoutes.GET("/check_username",controller.CheckUser)
	// To register a new user
	incommingRoutes.POST("/signup", do.Signup)
	// To login an existing user
	incommingRoutes.POST("/Login", do.Login)
}

//	Contains all the routes that are used do search an entity from the landing page

func Search(incommingRoutes *gin.Engine) {
	// To list the user,teams,tournaments,scrims (if any)
	// incommingRoutes.GET("/list_entity", controller.ListEntity)

	//To get the details of the details of the entity
	// incommingRoutes.GET("/seach_entity", controller.GetEntityDetails)

}

func ForgetPassword(incommingRoutes *gin.Engine) {
	// To get otp for the user do change the password
	// incommingRoutes.GET("/change_password", controller.ChangePassword)

	// To check the otp
	// incommingRoutes.GET("/user_otp", controller.ChangePasswordOtp)

	//	To change password
	// incommingRoutes.GET("/new_password", controller.ChangePassword)
}
