package routes

import (
	"bootcamp_es/controllers"
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type admin struct {
	adminController controllers.AdminControllers
	adminMw         middlewares.AdminCheckers
	mw              middlewares.Mwares
	auth            controllers.Auth
}

var ad admin

func Admin(incommingRoutes *gin.Engine) {
	admin := incommingRoutes.Group("/admin")
	admin.Use(ad.mw.AuthneticateToken, ad.adminMw.CheckUserType)
	admin.POST("/", ad.adminController.AdminHome)
	admin.GET("/searchcontent", ad.auth.SearchFirstFive)
	admin.GET("/search", ad.adminController.Search)
	admin.GET("/listuser/:page", ad.adminController.ListUsers)
	admin.GET("/listteam/:page", ad.adminController.ListTeam)
	admin.GET("/listtournament/:page", ad.adminController.ListTournament)
	admin.PUT("/updateusertype/:action", ad.adminController.UpdateUserType)
	admin.PUT("/updateuserblocked/:action", ad.adminController.UpdateBlock)
	admin.DELETE("/deletetournament/:tournament", ad.adminController.DeleteTournament)
	admin.DELETE("/deleteteam/:team", ad.adminController.DeleteTeam)
}
