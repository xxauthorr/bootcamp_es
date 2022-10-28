package database

import (
	"errors"
	"log"
)

type Check struct{}

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
