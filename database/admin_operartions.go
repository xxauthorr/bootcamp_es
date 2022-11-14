package database

import (
	"log"
)

type Admin struct {
}

func (a Admin) GetUserCount() string {
	var count string
	getStmnt := `SELECT count(id) FROM user_data;`
	res := Db.QueryRow(getStmnt)
	if err := res.Scan(&count); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return count
}
func (a Admin) GetTeamCount() string {
	var count string
	getStmnt := `SELECT count(id) FROM team_data;`
	res := Db.QueryRow(getStmnt)
	if err := res.Scan(&count); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return count
}
func (a Admin) GetTournamentCount() string {
	var count string
	getStmnt := `SELECT count(id) FROM tournament_data;`
	res := Db.QueryRow(getStmnt)
	if err := res.Scan(&count); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return count
}
