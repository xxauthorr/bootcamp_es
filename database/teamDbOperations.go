package database

import (
	"bootcamp_es/models"
	"fmt"
	"log"
	"time"
)

type Team struct {
}

type TeamProfileUpdate struct {
}

func (t Team) InsertTeamNotification(player, team, role string) error {
	time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	insertStmnt := `INSERT INTO user_notifications(team,player,role,time) VALUES ($1,$2,$3,$4)`
	_, err := Db.Exec(insertStmnt, team, player, role, time)
	if err != nil {
		return err
	}
	return nil
}

func (t Team) InsertTeam(team models.TeamReg, user string) error {
	insertStmnt := `INSERT INTO team_data(team_name,leader,instagram,discord) values ($1,$2,$3,$4);`
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

func (t Team) FetchTeamData(team string) {
	//  do it after the team edit
}

func (t TeamProfileUpdate) GetAchievmentsName(data models.TeamAchievementsAdd) string {
	var teamId, val string
	getTeamId := `SELECT id FROM team_data WHERE team_name = $1;`
	res := Db.QueryRow(getTeamId, data.TeamName)
	if err := res.Scan(&teamId); err != nil {
		fmt.Println(err.Error())
		log.Panic(err.Error())
		return ""
	}
	transaction.StartTransaction()
	insertStmnt := `INSERT INTO team_achievements (content,data,team_id) VALUES ($1,$2,$3);`
	_, err := Db.Exec(insertStmnt, data.Content, "sample", teamId)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	getStmnt := `SELECT id FROM user_achievements WHERE data = $1`
	row := Db.QueryRow(getStmnt, "sample")
	err = row.Scan(&val)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	return val
}

func (t TeamProfileUpdate) InsertTeamAchievements(location, id string) bool {
	updateStmnt := `UPDATE team_achievements SET data = $1 WHERE id = $2;`
	_, err := Db.Exec(updateStmnt, location, id)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return false
	}
	transaction.CommitTransaction()
	return true
}
