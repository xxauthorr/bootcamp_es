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
	//	To get otp for signup
	incommingRoutes.GET("/signup_otp", auth.SendPhoneOTP)
	//	To check the given otp from the user
	incommingRoutes.GET("/verify_otp", auth.CheckPhoneOtp)
	//	To check wheather the user exist or not
	incommingRoutes.GET("/verify_username", auth.CheckUser)
	//	To register a new user
	incommingRoutes.POST("/signup", auth.Signup)
	//	To login an existing user
	incommingRoutes.POST("/login", auth.Login)
	// 	To send otp for changing otp password
	incommingRoutes.GET("/forgot_password", auth.ForgotPassword)
	//	To check the otp for changing forget password
	incommingRoutes.GET("/verify_forgetotp", auth.VerifyPassOtp)
	//	To change the new password
	incommingRoutes.GET("/change_password", auth.ChangePassword)
}

//	Contains all the routes that are used auth search an entity from the landing page

func Search(incommingRoutes *gin.Engine) {
	// To list the user,teams,tournaments,scrims (if any)
	// incommingRoutes.GET("/list_entity", controller.ListEntity)

	//To get the details of the details of the entity
	// incommingRoutes.GET("/seach_entity", controller.GetEntityDetails)

}

func ForgetPassword(incommingRoutes *gin.Engine) {
	// To get otp for the user auth change the password
	// incommingRoutes.GET("/change_password", controller.ChangePassword)

	// To check the otp
	// incommingRoutes.GET("/user_otp", controller.ChangePasswordOtp)

	//	To change password
	// incommingRoutes.GET("/new_password", controller.ChangePassword)
}
