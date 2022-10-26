package main

import (
	routes "bootcamp_es/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	router := gin.New()

	routes.Authroutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
