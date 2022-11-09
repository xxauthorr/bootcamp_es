package database

import (
	bycrypt "bootcamp_es/services/byCrypt"
	"fmt"
	"log"
)

type Check struct {
	passHelper bycrypt.ByCrypt
}

type DBoperation struct{}

// if the given number exist return true
func (a Check) CheckPhoneNumber(number string) bool {
	checkStmt := `SELECT * FROM user_data WHERE phone = $1;`
	res, _ := Db.Exec(checkStmt, number)
	result, _ := res.RowsAffected()
	if result != 0 {
		return result != 0
	}
	return false
}

// if user exists return true
func (a Check) CheckUser(username string) bool {

	checkStmt := `SELECT * FROM user_data WHERE username = $1;`
	res, _ := Db.Exec(checkStmt, username)
	result, _ := res.RowsAffected()
	if result != 0 {
		return result != 0
	}
	return false
}

func (a Check) CheckTeam(teamName string) bool {
	checkStmt := `SELECT * FROM team_data WHERE team_name = $1;`
	res, _ := Db.Exec(checkStmt, teamName)
	result, _ := res.RowsAffected()
	if result != 0 {
		return result != 0
	}
	return false
}

func (a Check) CheckPassword(user, pass string) (bool, error) {
	var hash string
	getStmnt := `SELECT password FROM user_data WHERE username = $1;`
	rows := Db.QueryRow(getStmnt, user)
	if rows.Err() != nil {
		return false, rows.Err()
	}
	rows.Scan(&hash)
	res := a.passHelper.VerifyPassword(pass, hash)
	if !res {
		return res, nil
	}
	return res, nil
}

// returns true if user is in clan
func (a Check) CheckUserHasClan(user string) bool {
	checkStmnt := `SELECT * FROM team_data WHERE leader = $1;`
	res, _ := Db.Exec(checkStmnt, user)
	if res, _ := res.RowsAffected(); res != 0 {
		return true
	}
	return false
}

// returns the user type
func (a Check) CheckUserType(username string) string {
	var res string
	checkStmnt := `SELECT user_type FROM user_data WHERE username = $1;`
	row := Db.QueryRow(checkStmnt, username)
	if err := row.Scan(&res); err != nil {
		fmt.Println(err.Error())
		log.Panic(err.Error())
		return ""
	}
	return res
}

func (a Check) GetTeamFromLeader(leader string) string {
	var teamName string
	getTeam := `SELECT team_name FROM team_data WHERE leader = $1;`
	row := Db.QueryRow(getTeam, leader)
	if err := row.Scan(&teamName); err != nil {
		log.Panic(err.Error())
		return ""
	}
	return teamName
}

func (a DBoperation) StartTransaction() {
	_, err := Db.Exec(`BEGIN;`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (a DBoperation) RollBackTransaction() {
	_, err := Db.Exec(`ROLLBACK;`)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (a DBoperation) CommitTransaction() {
	_, err := Db.Exec(`COMMIT;`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func (a DBoperation) EndTransaction() {
	_, err := Db.Exec(`END;`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
