package database

import (
	"bootcamp_es/helpers"
	"errors"
	"log"
)

type Check struct {
	passHelper helpers.ByCrypt
}

func (a Check) CheckPhoneNumber(number string) error {
	checkStmt := `SELECT * FROM user_data WHERE phone = $1;`
	res, err := Db.Exec(checkStmt, number)
	if err != nil {
		log.Fatal(err)
		return err
	}
	result, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if result != 0 {
		return errors.New("true")
	}
	return nil
}

func (a Check) CheckUser(username string) error {

	checkStmt := `SELECT * FROM user_data WHERE username = $1;`
	res, err := Db.Exec(checkStmt, username)
	if err != nil {
		log.Fatal(err)
		return err
	}
	result, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if result != 0 {
		return errors.New("Exist")
	}
	return nil
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
