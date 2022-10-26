package routes

import(
	controller "bootcamp_es/controllers"
	"github.com/gin-gonic/gin")

func Authroutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/signup",controller.Signup)
	incommingRoutes.POST("/Login",controller.Login)
}
