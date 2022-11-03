package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserEdit struct {
}

type User struct {
	// check    database.Check
	// get      database.User
	// userData models.UserProfileData
}

func (user User) UserProfile(ctx *gin.Context) {
	// username := ctx.Param("username")
	// if username == "" {
	// 	ctx.JSON(http.StatusNotFound, nil)
	// 	return
	// }
	// status := user.check.CheckUser(username)
	// if !status {
	// 	ctx.JSON(http.StatusFound, gin.H{"status": status, "msg": "user not found"})
	// 	return
	// }
	// res, userData := user.get.UserProfileData(username)
	// if !res {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			ctx.JSON(http.StatusInternalServerError, err)
	// 		}
	// 	}()
	// }

}

func (edit UserEdit) ProfileEdit(ctx *gin.Context) {

}

func (edit UserEdit) UserAcheivementsEdit(ctx *gin.Context) {

}

func (edit UserEdit) UserSocialEdit(ctx *gin.Context) {

}

func (edit UserEdit) UserChangePassword(ctx *gin.Context) {

}
