package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xxauthorr/bootcamp_es/database"
	routes "github.com/xxauthorr/bootcamp_es/routes"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// loads env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		os.Exit(0)
	}
	port := os.Getenv("PORT")

	// connect to the psql database and return error if any
	if err := database.ConnectDb(); err != nil {
		log.Fatal(err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())

	routes.Authroutes(router)
	routes.UnAuthRoutes(router)
	routes.Team(router)
	routes.User(router)
	routes.Tournament(router)
	routes.Admin(router)
	router.NoRoute(routes.BadRoute)

	color.New(color.BgHiGreen).Print("server is running...")
	fmt.Println()
	router.Run(":" + port)
}
