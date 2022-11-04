package database

import "fmt"

type Get struct {
}

var transaction DBoperation

func (g Get) GetPhoneNumber(username string) string {
	var phone string
	getStmnt := `SELECT phone FROM user_data WHERE username = $1;`
	rows := Db.QueryRow(getStmnt, username)
	rows.Scan(&phone)
	return phone
}

func (g Get) GetUsername(phone string) string {
	var username string
	getStmnt := `SELECT username FROM user_data WHERE phone = $1;`
	rows := Db.QueryRow(getStmnt, phone)
	rows.Scan(&username)
	return username
}

// returns the count and if the user_data does'nt contain achievements id return false
func (g Get) GetNewAchievementName(user, content string) string {
	var id string
	var count string
	transaction.StartTransaction()
	getStmnt := `SELECT id FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getStmnt, user)
	if err := row.Scan(&id); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}

	countStmnt := `SELECT count(id) FROM user_achievements WHERE id = $1 AND type = $2;`
	value := Db.QueryRow(countStmnt,id,content)
	
	if err := value.Scan(&count); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	return count
}
