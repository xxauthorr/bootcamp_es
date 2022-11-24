package helpers

import (
	"log"
	"os"

	"github.com/xxauthorr/bootcamp_es/database"
	"github.com/xxauthorr/bootcamp_es/models"

	"github.com/joho/godotenv"
)

type AdminHelper struct {
	db       database.Admin
	entities models.Entities
	// check    database.Check
}

func (a AdminHelper) SuperUser(user string) bool {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		return false
	}
	superUser := os.Getenv("SUPER_USER")
	if superUser != user {
		return superUser == user
	}
	return true
}

func (a AdminHelper) GetEntitiesCount() models.Entities {
	count := a.db.GetUserCount()
	a.entities.Users = count
	count = a.db.GetTeamCount()
	a.entities.Teams = count
	count = a.db.GetTournamentCout()
	a.entities.Tournaments = count
	return a.entities
}
