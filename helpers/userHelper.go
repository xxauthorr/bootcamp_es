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

func (u UserHelper) FetchUserData(user string) models.UserProfileData {
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
