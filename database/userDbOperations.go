package database

import (
	model "bootcamp_es/models"
	bycrypt "bootcamp_es/services/byCrypt"
	"time"
)

type User struct {
	helper bycrypt.ByCrypt
}

func (u User) RegisterUser(user model.SignupForm) error {

	created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	*user.Password = u.helper.HashPassword(*user.Password)

	insertStm := `INSERT INTO user_data (username,phone,password,user_type,popularity,created_at,updated_at,block) VALUES ($1,$2,$3,'USER','0',$4,$5,'false');`
	_, err := Db.Exec(insertStm, *user.UserName, *user.Phone, *user.Password, created_at, updated_at)
	if err != nil {
		return err
	}
	return nil
}

func (u User) ChangePass(data, password string) error {
	password = u.helper.HashPassword(password)
	if data[0] == '+' {
		updateStmt := `UPDATE user_data SET password = $1 WHERE phone = $2;`
		_, err := Db.Exec(updateStmt, password, data)
		if err != nil {
			return err
		}
		return nil
	}
	updateStmt := `UPDATE user_data SET password = $1 WHERE username = $2;`
	_, err := Db.Exec(updateStmt, password, data)
	if err != nil {
		return err
	}
	return nil
}
