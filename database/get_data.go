package database

import (
	model "github.com/xxauthorr/bootcamp_es/models"
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

func (g Get) GetTournamentOwner(tour string) string {
	stmnt := `SELECT owner FROM tournament_data WHERE tournament_name = $1;`
	row := Db.QueryRow(stmnt, tour)
	if row.Err() != nil {
		fmt.Println(row.Err().Error())
		return ""
	}
	var owner string
	if err := row.Scan(&owner); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return owner
}

// returns the new id after inserting dummy data into user_data
func (g Get) GetNewAchievementName(user, content string) string {
	var userId, id string
	transaction.StartTransaction()
	getStmnt := `SELECT id FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getStmnt, user)
	if err := row.Scan(&userId); err != nil {
		transaction.RollBackTransaction()
		fmt.Println(err.Error())
		return ""
	}

	insertStmnt := `INSERT INTO user_achievements (content,data,user_id) VALUES ($1,$2,$3) RETURNING id;`
	val := Db.QueryRow(insertStmnt, content, "sample", userId)
	if val.Err() != nil {
		return ""
	}
	if err := val.Scan(&id); err != nil {
		return ""
	}
	return id
}

func (g Get) UserBio(user string) (
	bool, model.UserProfileData) {
	var userData model.UserProfileData
	getBioData := `SELECT id,username,phone,avatar,email,bio,team,crew,popularity,created_at FROM user_data WHERE username = $1;`
	row := Db.QueryRow(getBioData, user)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false, userData
	}
	err := row.Scan(&userData.Id, &userData.UserName, &userData.Phone, &userData.Avatar, &userData.Email, &userData.Bio, &userData.Team, &userData.Crew, &userData.Popularity, &userData.Created_at)
	if err != nil {
		log.Panic(err.Error())
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
	var notification model.U_Notification
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

func (g Get) CheckTeamCoLeader(user string) bool {
	var id int64
	stmnt := `SELECT count(id) FROM team_data WHERE co_leader = $1;`
	row := Db.QueryRow(stmnt, user)
	if row.Err() != nil {
		log.Fatal(row.Err().Error())
	}
	row.Scan(&id)
	if id != 0 {
		return id != 0
	}
	return false
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

func (g Get) CheckTeamExist(leader string) string {
	var teamName string
	getTeam := `SELECT team_name FROM team_data WHERE leader = $1;`
	row := Db.QueryRow(getTeam, leader)
	if err := row.Scan(&teamName); err != nil {
		// fmt.Println(err.Error(), "getTeamLeader")
		return ""
	}
	return teamName
}

func (g Get) GetTeamLeader(team string) string {
	var leader string
	stmnt := `SELECT leader FROM team_data WHERE team_name = $1;`
	row := Db.QueryRow(stmnt, team)
	if err := row.Scan(&leader); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return leader
}

func (g Get) GetTeamStrength(team string) int {
	var strength int
	row := Db.QueryRow(`SELECT count(id) FROM user_data WHERE team = $1`, team)
	if row.Err() != nil {
		fmt.Println(row.Err().Error())
		return 0
	}
	if err := row.Scan(&strength); err != nil {
		fmt.Println(row.Err().Error())
		return 0
	}
	return strength
}

func (g Get) TopEntities() model.HomeData {
	var data model.HomeData
	var player model.TopPlayers
	var team model.TopTeams
	stmnt := `SELECT username,popularity,team,avatar FROM user_data ORDER BY popularity ASC LIMIT 6;`
	rows, err := Db.Query(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&player.Player, &player.Popularity, &player.Team, &player.Avatar)
		if err != nil {
			fmt.Println(err.Error())
			return data
		}
		data.Top_Players = append(data.Top_Players, player)
	}
	stmnt = `SELECT team_name,leader,avatar FROM team_data ORDER BY team_name ASC LIMIT 6;`
	rows, err = Db.Query(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&team.Team, &team.Leader, &team.Avatar)
		if err != nil {
			fmt.Println(err.Error())
			return data
		}
		data.Top_teams = append(data.Top_teams, team)
	}
	return data
}

func (g Get) GetTournamentId(name string) int64 {
	var id int64
	row := Db.QueryRow(`SELECT id FROM tournament_data WHERE tournament_name = $1`, name)
	if row.Err() != nil {
		fmt.Println(row.Err().Error())
		return 0
	}
	if err := row.Scan(&id); err != nil {
		fmt.Println(row.Err().Error())
		return 0
	}
	return id
}
func (g Get) GetFirstFive(data model.Search) (bool, []string) {
	var value string
	var values []string
	data.Value = data.Value + "%"
	if data.Entity == "user" {
		stmnt := `SELECT username FROM user_data WHERE username LIKE $1 LIMIT 5;`
		rows, err := Db.Query(stmnt, data.Value)
		if err != nil {
			fmt.Println(err.Error())
			return false, values
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&value)
			values = append(values, value)
		}
		return true, values
	}
	if data.Entity == "team" {
		stmnt := `SELECT team_name FROM team_data WHERE team_name LIKE $1 LIMIT 5;`
		rows, err := Db.Query(stmnt, data.Value)
		if err != nil {
			fmt.Println(err.Error())
			return false, values
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&value)
			values = append(values, value)
		}
		return true, values
	}
	if data.Entity == "tournament" {
		stmnt := `SELECT tournament_name FROM tournament_data WHERE tournament_name LIKE $1 LIMIT 5;`
		rows, err := Db.Query(stmnt, data.Value)
		if err != nil {
			fmt.Println(err.Error())
			return false, values
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&value)
			values = append(values, value)
		}
		return true, values
	}
	// if data.Entity == "scrims" {

	// }
	return false, values
}
