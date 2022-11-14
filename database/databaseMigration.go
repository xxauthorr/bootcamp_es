package database

import (
	"bootcamp_es/models"
	"fmt"

	"gorm.io/gorm"
)

type Models struct {
	user_data         models.User_data
	user_achievement  models.User_achievement
	user_notification models.User_notification
	user_social       models.User_social
	user_popularity   models.User_popularity
	team_data         models.Team_data
	team_achievement  models.Team_achievement
	team_notification models.Team_notification
	tournament_data   models.Tournament_data
}

var migrate Models

func AutoMigrateTables(mig *gorm.DB) {
	if err := mig.AutoMigrate(&migrate.user_data); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.team_data); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.user_achievement); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.user_notification); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.user_social); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.team_data); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.team_achievement); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.team_notification); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.user_popularity); err != nil {
		fmt.Println(err.Error())
	}
	if err := mig.AutoMigrate(&migrate.tournament_data); err != nil {
		fmt.Println(err.Error())
	}
}
