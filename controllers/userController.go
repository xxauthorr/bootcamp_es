package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/middlewares"
	"bootcamp_es/models"
	amazons3 "bootcamp_es/services/AmazonS3"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserEdit struct {
	check               database.Check
	transaction         database.DBoperation
	userBioData         models.UserBioEdit
	userAddAchievements models.UserAchievementsAdd
	userDelAchievements models.UserAchievementsDel
	userSocial          models.UserSocialEdit
	s3                  amazons3.S3
	update              database.User
	achievement         database.Get
}

type User struct {
	check database.Check
	// get   database.User
	help  helpers.User
}

func (user User) UserProfile(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	status := user.check.CheckUser(username)
	if !status {
		ctx.JSON(http.StatusFound, gin.H{"status": status, "msg": "user not found"})
		return
	}
	userData := user.help.FetchUserData(username)
	ctx.JSON(http.StatusOK, gin.H{"user": userData})

}

func (edit UserEdit) BioEdit(ctx *gin.Context) {
	user := middlewares.TokenUser
	if user == "" {
		fmt.Println("user is empty in bioedit ")
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	if res := edit.check.CheckUser(user); !res {
		ctx.JSON(http.StatusForbidden, res)
		return
	}
	if err := ctx.ShouldBind(&edit.userBioData); err != nil {
		fmt.Println("error in bson bind")
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	edit.userBioData.UserName = user
	if err := validate.Struct(edit.userBioData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	avatar, err := edit.s3.UploadToS3(edit.userBioData.Avatar, user+"BootCamp.jpg")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	if err := edit.update.UpdateBio(edit.userBioData, avatar); err != nil {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	//  then go to /profile/:username
	ctx.JSON(http.StatusOK, true)
}

func (edit UserEdit) UserAcheivementsAdd(ctx *gin.Context) {
	user := middlewares.TokenUser
	if user == "" {
		fmt.Println("user is empty in bioedit ")
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	if res := edit.check.CheckUser(user); !res {
		ctx.JSON(http.StatusForbidden, res)
		return
	}
	if err := ctx.ShouldBind(&edit.userAddAchievements); err != nil {
		fmt.Println("bind error")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	edit.userAddAchievements.UserName = user
	if err := validate.Struct(edit.userAddAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	val := edit.achievement.GetNewAchievementName(user, edit.userAddAchievements.Content)
	if val == "" {
		fmt.Println("no value")
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	location, err := edit.s3.UploadToS3(edit.userAddAchievements.Data, user+"_"+edit.userAddAchievements.Content+"_"+val+".jpg")
	if err != nil {
		edit.transaction.RollBackTransaction()
		fmt.Println("s3 error")
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	res := edit.update.InsertAchievements(val, location)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "insert error"})
		return
	}
	ctx.JSON(http.StatusOK, true)
}

func (edit UserEdit) UserAcheivementsDelete(ctx *gin.Context) {
	user := middlewares.TokenUser
	if err := ctx.BindJSON(&edit.userDelAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	edit.userDelAchievements.UserName = user
	if err := validate.Struct(edit.userDelAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res := edit.check.CheckUser(user); !res {
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	res := edit.update.DeleteAchievement(edit.userDelAchievements.Data)
	if !res {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	ctx.JSON(http.StatusOK, true)
}

func (edit UserEdit) UserSocialEdit(ctx *gin.Context) {
	user := middlewares.TokenUser
	if err := ctx.BindJSON(&edit.userSocial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	edit.userSocial.UserName = user
	if err := validate.Struct(edit.userSocial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := edit.update.UserSocialUpdate(edit.userSocial)
	if res != "true" {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	ctx.JSON(http.StatusOK, true)

}

func (edit UserEdit) UserChangePassword(ctx *gin.Context) {

}
