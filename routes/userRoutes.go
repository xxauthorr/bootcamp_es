package routes

import (
	"bootcamp_es/middlewares"

	"github.com/gin-gonic/gin"
)

type route struct {
	mw middlewares.Mwares
}

// Contains all the routes that are used for the user profile (not to edit)

func User(incommingRoutes *gin.Engine) {
	var route route
	incommingRoutes.Use(route.mw.Authneticate)

}

// Contains all the routes to edit the control settings (password,email,phone)
func UserSettigs(incommingRoutes *gin.Engine) {

}
