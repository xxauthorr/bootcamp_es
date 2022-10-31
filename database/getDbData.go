package database

type Get struct {
}

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
