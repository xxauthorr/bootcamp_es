package main

import (
	"bootcamp_es/database"
	routes "bootcamp_es/routes"
	"log"
	"os"

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
	database.ConnectDb()

	router := gin.New()
	router.Use(gin.Logger())

	routes.Authroutes(router)
	routes.Team(router)
	routes.User(router)
	routes.Tournament(router)
	routes.Admin(router)

	router.Run(":" + port)
}
