package routes

import (
	"bootcamp_es/controllers"

	"github.com/gin-gonic/gin"
)

var (
	adminController controllers.AdminControllers
)

func Admin(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(mw.AuthneticateToken, checker.CheckUserType)
	admin := incommingRoutes.Group("/admin")
	admin.POST("/search", adminController.Search)

}
