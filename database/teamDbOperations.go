package database

import (
	"bootcamp_es/models"
	"fmt"
	"time"
)

type Team struct {
	models.TeamReg
}

func (t *Team) InsertTeamNotification(player, team,role string) error {
	time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	insertStmnt := `INSERT INTO notification(team,player,time,role) VALUES ($1,$2,$3,$4)`
	_, err := Db.Exec(insertStmnt, team, player, time,role)
	if err != nil {
		return err
	}
	return nil
}

func (t *Team) RegisterTeam(team models.TeamReg, leader string) error {
	fmt.Println("entered")
	insertStmnt := `INSERT INTO team_data(team_name,leader,instagram,discord) values ($1,$2,$3,$4);`
	_, err := Db.Exec(insertStmnt, *team.TeamName, leader, *team.Instagram, *team.Discord)
	if err != nil {
		return err
	}
	return nil
}
