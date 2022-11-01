package database

import (
	"bootcamp_es/models"
)

type Insert struct {
	models.TeamReg
}

func (t *Insert) InsertTeamNotification(player, team, role string) error {

	insertStmnt := `INSERT INTO user_notify(team,player,role) VALUES ($1,$2,$3)`
	_, err := Db.Exec(insertStmnt, team, player, role)
	if err != nil {
		return err
	}
	return nil
}

func (t *Insert) InsertTeam(team models.TeamReg, user string) error {
	insertStmnt := `INSERT INTO team_data(team_name,leader,istagram,discord) values ($1,$2,$3,$4);`
	_, err := Db.Exec(insertStmnt, team.TeamName, user, team.Instagram, team.Discord)
	if err != nil {
		return err
	}
	insertStmnt = `UPDATE user_data SET team = $1 WHERE username = $2;`
	_, err = Db.Exec(insertStmnt, team.TeamName, user)
	if err != nil {
		return err
	}
	return nil
}
