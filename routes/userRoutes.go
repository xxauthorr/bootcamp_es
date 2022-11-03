package routes

import (
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

var mw middlewares.Mwares

// Contains all the routes that are used for the user profile (not to edit)

func User(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(mw.AuthneticateToken)
	incommingRoutes.Group("/profile")
	incommingRoutes.GET("/:username",)
	incommingRoutes.GET("/user_get", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "nothing"})
	})
}

// Contains all the routes to edit the control settings (password,email,phone)
// func UserSettings(incommingRoutes *gin.Engine) {

// }
