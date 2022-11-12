package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
)

type AdminHelper struct {
	db       database.Admin
	entities models.Entities
	// check    database.Check
}

func (a AdminHelper) GetEntitiesCount() models.Entities {
	count := a.db.GetUserCount()
	a.entities.Users = count
	count = a.db.GetTeamCount()
	a.entities.Teams = count
	return a.entities
}

// func (a AdminHelper) AdminSearch(search models.Search) (bool, models.Search) {
// 	if search.Entity == "user" {
// 		if res := a.check.CheckUser(search.Value); !res {
// 			return false, search
// 		}
// 	}
// 	if search.Entity == "team" {

// 	}
// 	return true, search
// }
