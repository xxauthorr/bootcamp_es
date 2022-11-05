package database

import (
	model "bootcamp_es/models"
	bycrypt "bootcamp_es/services/byCrypt"
	"fmt"
	"time"
)

type User struct {
	helper bycrypt.ByCrypt
}

func (u User) InsertUser(user model.SignupForm) error {

	created_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Password = u.helper.HashPassword(user.Password)

	insertStm := `INSERT INTO user_data (username,phone,password,user_type,popularity,created_at,updated_at,block) VALUES ($1,$2,$3,'USER','0',$4,$5,'false');`
	_, err := Db.Exec(insertStm, user.UserName, user.Phone, user.Password, created_at, updated_at)
	if err != nil {
		return err
	}
	return nil
}

func (u User) ChangePass(data, password string) error {
	password = u.helper.HashPassword(password)
	if string(data[0]) == "+" {
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

func (u User) UpdateBio(data model.UserBioEdit, avatar string) error {
	updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateStmnt := "UPDATE user_data SET bio = $1,crew = $2,role = $3,avatar = $4,updated_at = $5 WHERE username = $6;"
	_, err := Db.Exec(updateStmnt, data.Bio, data.Crew, data.Role, avatar, updated_at, data.UserName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (u User) InsertAchievements(id, location string) bool {
	updateStmnt := `UPDATE user_achievements SET data = $1 WHERE id = $2;`
	_, err := Db.Exec(updateStmnt, location, id)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return false
	}
	transaction.CommitTransaction()
	return true
}

func (u User) DeleteAchievement(data string) bool {
	delStmnt := `DELETE FROM user_achievements WHERE data = $1;`
	_, err := Db.Exec(delStmnt, data)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (u User) UserSocialUpdate(user model.UserSocialEdit) string {
	transaction.StartTransaction()
	var userId, socailId int64
	socailId = 0
	getStmnt := `SELECT id FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getStmnt, user.UserName)
	if err := row.Scan(&userId); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	CheckStmnt := `SELECT id FROM user_social WHERE id = $1;`
	row = Db.QueryRow(CheckStmnt, userId)
	if err := row.Scan(&socailId); err != nil {
		fmt.Println(socailId)
		if socailId == 0 {
			fmt.Println("working")
			insertStmnt := `INSERT INTO user_social(id,instagram,whatsapp,discord) VALUES ($1,$2,$3,$4);`
			_, err := Db.Exec(insertStmnt, userId, user.Instagram, user.Whatsapp, user.Discord)
			if err != nil {
				fmt.Println("error is here ")
				transaction.RollBackTransaction()
				fmt.Println(err.Error())
				return ""
			}
			transaction.CommitTransaction()
			return "true"
		}
	}
	updateStmnt := `UPDATE user_social SET instagram = $1,whatsapp = $2,discord=$3 WHERE id = $4;`
	_, err := Db.Exec(updateStmnt, user.Instagram, user.Whatsapp, user.Discord, userId)
	if err != nil {
		fmt.Println("lsjdhfl")
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	transaction.CommitTransaction()
	return "true"

}


