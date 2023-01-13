package helpers

import (
	"fmt"

	"github.com/xxauthorr/bootcamp_es/database"
	"github.com/xxauthorr/bootcamp_es/models"
)

type UserHelper struct {
	get   database.Get
	user  database.User
	check database.Check
}

func (u UserHelper) FetchProfileData(user string, status bool) models.UserProfileData {
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
	userData.Visit = false
	if status {
		userData.Visit = true
		res, userData = u.get.UserNotification(userData)
		if !res {
			fmt.Println("error notificatiion")
			return userData
		}
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

func (u UserHelper) UpdatePopularity(data models.UserPopularityUpdate) bool {
	if data.Action == "true" {
		res := u.check.CheckPopularityList(data.From, data.To)
		if !res {
			return true
		}
	}
	if data.Action == "false" {
		res := u.check.CheckPopularityList(data.From, data.To)
		if res {
			return true
		}
	}
	if res := u.user.UpdatePopularity(data.To, data.Action); !res {
		fmt.Println("err 1")
		return res
	}
	if res := u.user.UpdatePopularityList(data); !res {
		fmt.Println("err 2")

		return res
	}
	return true
}
