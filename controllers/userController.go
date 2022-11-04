package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/middlewares"
	"bootcamp_es/models"
	amazons3 "bootcamp_es/services/AmazonS3"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserEdit struct {
	check               database.Check
	userBioData         models.UserBioEdit
	userAchievements    models.UserAchievementsEdit
	s3                  amazons3.S3
	update              database.User
	addAchievement database.Get
}

type User struct {
	check database.Check
	get   database.User
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
	res, _ := user.get.UserProfileData(username)
	if !res {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			}
		}()
	}

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
	if err := ctx.ShouldBind(&edit.userAchievements); err != nil {
		fmt.Println("bind error")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	edit.userAchievements.UserName = user
	if err := validate.Struct(edit.userAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count := edit.addAchievement.GetNewAchievementName(user, edit.userAchievements.Content)
	location, err := edit.s3.UploadToS3(edit.userAchievements.Achievement, user+"_"+edit.userAchievements.Content+"_"+count+".jpg")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	res := edit.update.InsertAchievements(edit.userAchievements.UserName,edit.userAchievements.Content,location)
	if !res {
		ctx.JSON(http.StatusInternalServerError,gin.H{"msg":"insert error"})
		return
	}
	ctx.JSON(http.StatusOK, true)
}

func (edit UserEdit) UserAcheivementsDelete(ctx *gin.Context) {

}

func (edit UserEdit) UserSocialEdit(ctx *gin.Context) {

}

func (edit UserEdit) UserChangePassword(ctx *gin.Context) {

}
