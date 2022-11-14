package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type admin struct {
	adminController controllers.AdminControllers
	adminMw           middlewares.AdminCheckers
	mw    middlewares.Mwares

}

var ad admin

func Admin(incommingRoutes *gin.Engine) {
	admin := incommingRoutes.Group("/admin")
	admin.Use(ad.mw.AuthneticateToken)
	admin.Use(ad.adminMw.CheckUserType)
	admin.POST("/search", ad.adminController.Search)

}
