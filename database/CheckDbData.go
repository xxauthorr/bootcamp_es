package database

import (
	bycrypt "bootcamp_es/services/byCrypt"
	"errors"
	"log"
)

type Check struct {
	passHelper bycrypt.ByCrypt
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
		return errors.New("Exist")
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

func (a Check) TeamLeaderCheck(leader string) (bool, error) {
	checkStmnt := `SELECT * FROM team_data WHERE leader = $1;`
	res, err := Db.Exec(checkStmnt, leader)
	if err != nil {
		return false, err
	}
	if res, _ := res.RowsAffected(); res != 0 {
		return false, nil
	}
	return true, nil
}
