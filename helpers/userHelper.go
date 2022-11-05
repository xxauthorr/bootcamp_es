package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
	"fmt"
)

type User struct {
	get database.Get
}

func (u User) FetchUserData(user string) models.UserProfileData {
	res, userData := u.get.UserBio(user)
	if !res {
		fmt.Println("error")
		return userData
	}
	res, userData = u.get.UserSocial(userData)
	if !res {
		fmt.Println("error")
		return userData
	}

	res, userData = u.get.UserAchievements(userData)
	if !res {
		fmt.Println("error")
		return userData
	}
	res, userData = u.get.UserNotification(userData)
	if !res {
		fmt.Println("error")
		return userData
	}
	return userData
}
