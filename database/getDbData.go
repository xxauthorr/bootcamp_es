package database

import (
	model "bootcamp_es/models"
	"fmt"
	"log"
)

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

// returns the new id after inserting dummy data into user_data
func (g Get) GetNewAchievementName(user, content string) string {
	var userId, val string
	transaction.StartTransaction()
	getStmnt := `SELECT id FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getStmnt, user)
	if err := row.Scan(&userId); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}

	insertStmnt := `INSERT INTO user_achievements (content,data,user_id) VALUES ($1,$2,$3);`
	_, err := Db.Exec(insertStmnt, content, "sample", userId)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	getStmnt = `SELECT id FROM user_achievements WHERE data = $1`
	row = Db.QueryRow(getStmnt, "sample")
	err = row.Scan(&val)
	if err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}
	return val
}

func (g Get) UserBio(user string) (
	bool, model.UserProfileData) {
	var userData model.UserProfileData
	getBioData := `SELECT id,username,phone,email,bio,team,crew,popularity,created_at FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getBioData, user)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false, userData
	}
	err := row.Scan(&userData.Id, &userData.UserName, &userData.Phone, &userData.Email, &userData.Bio, &userData.Team, &userData.Crew, &userData.Popularity, &userData.Created_at)
	if err != nil {
		fmt.Println(err.Error())
	}
	return true, userData
}

func (g Get) UserSocial(userData model.UserProfileData) (bool, model.UserProfileData) {

	res, err := Db.Exec(`SELECT * FROM user_socials WHERE id = $1;`, userData.Id)
	if err != nil {
		fmt.Println(err.Error(), "social check")
		return false, userData
	}
	result, _ := res.RowsAffected()
	if result != 0 {
		getSocialData := `SELECT instagram,whatsapp,discord FROM user_socials WHERE Id = $1;`
		row := Db.QueryRow(getSocialData, userData.Id)
		if row.Err() != nil {
			log.Panic(row.Err())
			return false, userData
		}
		if err := row.Scan(&userData.Instagram, &userData.Whatsapp, &userData.Discord); err != nil {
			fmt.Println(err.Error())
			return false, userData
		}
	}
	return true, userData
}

func (g Get) UserAchievements(data model.UserProfileData) (bool, model.UserProfileData) {
	var ingame, tourney *string
	res, err := Db.Exec(`SELECT * FROM user_achievements WHERE user_id = $1 AND content = 'INGAME';`, data.Id)
	if err != nil {
		return false, data
	}
	count, _ := res.RowsAffected()
	if count != 0 {
		getIngame := `SELECT data FROM user_achievements WHERE user_id = $1 AND content = 'INGAME';`
		rows, err := Db.Query(getIngame, data.Id)
		if err != nil {
			fmt.Println(err.Error())
			return false, data
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&ingame); err != nil {
				fmt.Println(err.Error())
				return false, data
			}
			data.UserAchievements.Ingame = append(data.UserAchievements.Ingame, *ingame)
		}
	}
	res, err = Db.Exec(`SELECT * FROM user_achievements WHERE user_id = $1 AND content = 'TOURNAMENT';`, data.Id)
	if err != nil {
		return false, data
	}
	count, _ = res.RowsAffected()
	if count != 0 {
		getIngame := `SELECT data FROM user_achievements WHERE user_id = $1 AND content = 'TOURNAMENT';`
		rows, err := Db.Query(getIngame, data.Id)
		if err != nil {
			fmt.Println(err.Error())
			return false, data
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&tourney); err != nil {
				fmt.Println(err.Error())
				return false, data
			}
			data.UserAchievements.Tourney = append(data.UserAchievements.Tourney, *tourney)
		}
	}
	return true, data
}

func (g Get) UserNotification(data model.UserProfileData) (bool, model.UserProfileData) {
	var notification model.Notification
	getNotification := `SELECT id,team,role,time FROM user_notifications WHERE player = $1;`
	rows, err := Db.Query(getNotification, data.UserName)
	if err != nil {
		fmt.Println(err.Error())
		return false, data
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&notification.Id, &notification.Team, &notification.Role, &notification.Time); err != nil {
			fmt.Println(err.Error())
			return false, data
		}
		data.UserNotifications = append(data.UserNotifications, notification)
	}
	return true, data
}

func (g Get) GetTeamName(data string) string {
	var team string
	getTeam := `SELECT team FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getTeam, data)
	if err := row.Scan(&team); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return team
}
