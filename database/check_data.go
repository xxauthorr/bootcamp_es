package database

import (
	bycrypt "bootcamp_es/services/byCrypt"
	"fmt"
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

func (a Check) CheckTournament(tournament string) bool {
	checkStmt := `SELECT * FROM tournament_data WHERE tournament_name = $1;`
	res, _ := Db.Exec(checkStmt, tournament)
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
	var team string
	checkStmnt := `SELECT team FROM user_data WHERE username = $1;`
	res := Db.QueryRow(checkStmnt, user)
	if res.Err() != nil {
		// panic
		fmt.Println(res.Err().Error())
		return false
	}
	res.Scan(&team)
	if team != "" {
		return team != ""
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
		return ""
	}
	return res
}

func (a Check) CheckUserPopularity(from, to string, liked chan bool) {
	var count string
	stmnt := `SELECT count(id) FROM user_popularities WHERE provide = $1 AND consume = $2;`
	row := Db.QueryRow(stmnt, from, to)
	if err := row.Scan(&count); err != nil {
		fmt.Println(err.Error(), from, to)
		fmt.Println("error in checkuserPopularity")
		liked <- false
		return
	}
	if count != "1" {
		liked <- false
		return
	}
	liked <- true
}

func (a Check) CheckPopularityList(from, to string) bool {
	var val int
	stmnt := `SELECT count(id) FROM user_popularities WHERE provide = $1 AND consume = $2;`
	row := Db.QueryRow(stmnt, from, to)
	if err := row.Scan(&val); err != nil {
		fmt.Println(err.Error())
		return false
	}
	if val == 0 {
		return true
	}
	return false
}

func (a Check) CheckUserBlocked(user string) bool {
	var value string
	stmnt := `SELECT block FROM user_data WHERE username = $1;`
	row := Db.QueryRow(stmnt, user)
	if row.Err() != nil {
		fmt.Println(row.Err().Error())
		return false
	}
	if err := row.Scan(&value); err != nil {
		fmt.Println(err.Error())
		return false
	}
	if value == "false" {
		return value == "false"
	}
	return false
}

func (a Check) GetTour(user string) (string, bool) {
	var tourName *string
	stmnt := `SELECT tournament_name FROM tournament_data WHERE owner = $1;`
	row := Db.QueryRow(stmnt, user)
	if row.Err() != nil {
		fmt.Println(row.Err().Error())
		return "", false
	}
	if err := row.Scan(&tourName); err != nil {
		fmt.Println(row.Err().Error())
		return "", false
	}
	if tourName == nil {
		return "", false
	}
	return *tourName, true
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
