package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
	"fmt"
)

type UserHelper struct {
	get  database.Get
	user database.User
}

func (u UserHelper) FetchProfileData(user string) models.UserProfileData {
	res, userData := u.get.UserBio(user)
	if !res {
		fmt.Println("error bio")
		return userData
	}
	res, userData = u.get.UserSocial(userData)
	if !res {
		fmt.Println("error social")
		return userData
	}

	res, userData = u.get.UserAchievements(userData)
	if !res {
		fmt.Println("error achievement")
		return userData
	}
	res, userData = u.get.UserNotification(userData)
	if !res {
		fmt.Println("error notificatiion")
		return userData
	}
	return userData
}

func (u UserHelper) FetchUserData(user string) models.UserData {
	var data models.UserData
	res, userData := u.get.UserBio(user)
	if !res {
		fmt.Println("error bio")
		return data
	}
	res, userData = u.get.UserSocial(userData)
	if !res {
		fmt.Println("error social")
		return data
	}

	res, userData = u.get.UserAchievements(userData)
	if !res {
		fmt.Println("error achievement")
		return data
	}
	data.UserName = userData.UserName
	data.Avatar = userData.Avatar
	data.Team = userData.Team
	data.Crew = userData.Crew
	data.Phone = userData.Phone
	data.Bio = userData.Bio
	data.Discord = userData.Discord
	data.Instagram = userData.Instagram
	data.Email = userData.Email
	data.Created_at = userData.Created_at
	data.Popularity = userData.Popularity
	for i := range userData.UserAchievements.Ingame {
		data.UserAchievements.Ingame = append(data.UserAchievements.Ingame, userData.UserAchievements.Ingame[i])
	}
	for i := range userData.UserAchievements.Tourney {
		data.UserAchievements.Ingame = append(data.UserAchievements.Ingame, userData.UserAchievements.Tourney[i])
	}
	return data
}

func (u UserHelper) UpdateNotification(id, action string) bool {
	if action == "true" {
		if res := u.user.UpdateNotification(id); !res {
			return false
		}
		return true
	}
	if res := u.user.DelNotification(id); !res {
		return false
	}
	return true
}

func (u UserHelper) UpdatePopularity(data models.UserPopularityUpdate) bool {
	if res := u.user.AddPopularity(data.To, data.Action); !res {
		return res
	}
	if res := u.user.UpdatePopularityList(data); !res {
		return res
	}
	return true
}
