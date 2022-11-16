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

func (t TeamProfileUpdate) InsertTeamNotification(player, team, role string) error {
	time := time.Now().Format("2-01-2006 3:04:05 PM")

	insertStmnt := `INSERT INTO user_notifications(team,player,role,time) VALUES ($1,$2,$3,$4)`
	_, err := Db.Exec(insertStmnt, team, player, role, time)
	if err != nil {
		return err
	}
	return nil
}

func (t Team) InsertTeam(team models.TeamReg, user string) error {
	created_at := time.Now().Format("2-01-2006 3:04:05 PM")
	insertStmnt := `INSERT INTO team_data(team_name,leader,instagram,discord,created_at) values ($1,$2,$3,$4,$5);`
	_, err := Db.Exec(insertStmnt, team.TeamName, user, team.Instagram, team.Discord, created_at)
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

func (t Team) FetchTeamData(team string) models.TeamData {
	//  do it after the team edit
	var data models.TeamData
	var (
		value, content string
		id             int64
	)
	stmnt := `SELECT id,team_name,leader,bio,instagram,discord,youtube,avatar,co_leader FROM team_data WHERE team_name = $1;`
	row := Db.QueryRow(stmnt, team)
	if err := row.Scan(&id, &data.TeamName, &data.Leader, &data.Bio, &data.Instagram, &data.Discord, &data.Youtube, &data.Avatar, &data.Co_leader); err != nil {
		fmt.Println(err.Error())
		return data
	}
	stmnt = `SELECT content,data FROM team_achievements WHERE team_id = $1;`
	rows, err := Db.Query(stmnt, id)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer rows.Close()
	for rows.Next() {
		if err := row.Scan(&content, &value); err != nil {
			fmt.Println(err.Error())
			return data
		}
		if content == "TOURNAMENT" {
			data.Achievements.Tournaments = append(data.Achievements.Tournaments, value)
		} else {
			data.Achievements.Scrims = append(data.Achievements.Scrims, value)
		}
	}
	return data
}

func (t Team) FetchTeamNotification(data models.TeamData) models.TeamData {
	var value models.TeamNotification
	stmnt := `SELECT id,time,request,player FROM team_notifications WHERE team = $1;`
	rows, err := Db.Query(stmnt, data.TeamName)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&value.Id, &value.Time, &value.Request, &value.Player); err != nil {
			fmt.Println(err.Error())
			return data
		}
		data.Notifications = append(data.Notifications, value)
	}
	return data
}

func (t TeamProfileUpdate) GetAchievmentsName(data models.TeamAchievementsAdd) string {
	var teamId, id string
	transaction.StartTransaction()
	getTeamId := `SELECT id FROM team_data WHERE team_name = $1;`
	res := Db.QueryRow(getTeamId, data.TeamName)
	if err := res.Scan(&teamId); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		log.Panic(err.Error())
		return ""
	}
	// transaction.StartTransaction()
	insertStmnt := `INSERT INTO team_achievements (content,data,team_id) VALUES ($1,$2,$3) RETURNING id;`
	val := Db.QueryRow(insertStmnt, data.Content, "sample", teamId)
	if val.Err() != nil {
		transaction.RollBackTransaction()
		return ""
	}
	if err := val.Scan(&id); err != nil {
		transaction.RollBackTransaction()
		return ""
	}
	return id
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

func (t TeamProfileUpdate) DeleteTeamAchievements(data string) bool {
	fmt.Println(data)
	delStmnt := `DELETE FROM team_achievements WHERE data=$1;`
	if _, err := Db.Exec(delStmnt, data); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (t TeamProfileUpdate) UpdateBio(data models.TeamBioEdit, location string) bool {
	var err error
	if location != "" {
		insertStmnt := `UPDATE team_data SET instagram = $1,discord=$2,youtube=$3,bio=$4,avatar=$5 WHERE team_name = $6;`
		_, err = Db.Exec(insertStmnt, data.Instagram, data.Youtube, data.Discord, data.Bio, location, data.TeamName)
	} else {
		insertStmnt := `UPDATE team_data SET instagram = $1,discord=$2,youtube=$3,bio=$4 WHERE team_name = $5;`
		_, err = Db.Exec(insertStmnt, data.Instagram, data.Youtube, data.Discord, data.Bio, data.TeamName)
	}
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (t TeamProfileUpdate) DeleteCoLeader(user string) {
	stmnt := `UPDATE team_data SET co_leader = null WHERE co_leader = $1;`
	_, err := Db.Exec(stmnt, user)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (t TeamProfileUpdate) UpdateTeamNotification(data models.Notification, team string) bool {
	if data.Action == "false" {
		stmnt := `DELETE FROM team_notifications WHERE id = $1;`
		if _, err := Db.Exec(stmnt, data.Id); err != nil {
			fmt.Println(err.Error())
			return false
		}
		return true
	}
	if data.Action == "true" {
		var user string
		transaction.StartTransaction()
		getUser := `SELECT player FROM team_notifications WHERE id = $1;`
		row := Db.QueryRow(getUser, data.Id)
		if row.Err() != nil {
			transaction.RollBackTransaction()
			fmt.Println(row.Err().Error())
			return false
		}
		row.Scan(&user)
		// update user team
		updateStmnt := `UPDATE user_data SET team = $1 WHERE username = $2;`
		_, err := Db.Exec(updateStmnt, team, user)
		if err != nil {
			transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false
		}
		stmnt := `DELETE FROM team_notifications WHERE id = $1;`
		if _, err := Db.Exec(stmnt, data.Id); err != nil {
			transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false
		}
		transaction.CommitTransaction()
		return true
	}
	return false
}
