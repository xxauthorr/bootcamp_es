package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"
	amazons3 "bootcamp_es/services/AmazonS3"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserEdit struct {
	check               database.Check
	transaction         database.DBoperation
	update              database.User
	achievement         database.Get
	userBioData         models.UserBioEdit
	userAddAchievements models.UserAchievementsAdd
	userDelAchievements models.UserAchievementsDel
	userSocial          models.UserSocialEdit
	userNotification    models.Notification
	popularityUpdate    models.UserPopularityUpdate
	help                helpers.UserHelper
	s3                  amazons3.S3
}

type User struct {
	check      database.Check
	help       helpers.UserHelper
	getHelp    helpers.Help
	auth       helpers.Help
	AuthResult models.AuthResult
}

func (user User) UserProfile(ctx *gin.Context) {
	username := ctx.Param("username")
	status := user.check.CheckUser(username)
	if !status {
		ctx.JSON(http.StatusFound, gin.H{"status": status, "message": "User not found", "result": nil})
		return
	}
	res := user.auth.Authorize(ctx)
	client := ctx.GetString("user")
	if username == client && res {
		data := user.help.FetchProfileData(client, true)
		data.Liked = user.check.CheckUserPopularity(client, username)
		data.Visit = false
		user.AuthResult.User = data
		//update token
		token := user.getHelp.GetToken(client)
		user.AuthResult.Authorization = token
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Succesfully completed", "result": user.AuthResult})
		return
	}
	data := user.help.FetchProfileData(username, false)
	if !res {
		data.Visit = true
		data.Liked = false
		user.AuthResult.User = data
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Succesfully completed", "result": user.AuthResult})
		return
	}
	data.Liked = user.check.CheckUserPopularity(client, username)
	data.Visit = true
	user.AuthResult.User = data
	//update token
	token := user.getHelp.GetToken(client)
	user.AuthResult.Authorization = token
	//update token
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Succesfully completed", "result": user.AuthResult})
}

func (edit UserEdit) UserPopularityEdit(ctx *gin.Context) {
	user := ctx.GetString("user")
	to := ctx.Param("to")
	if res := edit.check.CheckUser(to); !res {
		ctx.JSON(http.StatusSeeOther, gin.H{"status": false, "message": "Invalid params", "result": nil})
		return
	}
	if res := edit.check.CheckUser(user); !res {
		ctx.JSON(http.StatusSeeOther, gin.H{"status": false, "message": "Invalid token claims", "result": nil})
		return
	}
	if err := ctx.BindJSON(&edit.popularityUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !"})
		return
	}
	edit.popularityUpdate.To = to
	edit.popularityUpdate.From = user
	if err := validate.Struct(edit.popularityUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !"})
		return
	}
	if res := edit.help.UpdatePopularity(edit.popularityUpdate); !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request successfully completed"})
}

func (edit UserEdit) BioEdit(ctx *gin.Context) {
	user := ctx.GetString("user")
	if res := edit.check.CheckUser(user); !res {
		ctx.JSON(http.StatusSeeOther, gin.H{"status": false, "message": "Invalid token claims", "result": nil})
		return
	}
	if err := ctx.ShouldBind(&edit.userBioData); err != nil {
		fmt.Println("error in bson bind :", err.Error())
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	edit.userBioData.UserName = user
	if err := validate.Struct(edit.userBioData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !"})
		return
	}

	if _, ProfilePicture, _ := ctx.Request.FormFile("avatar"); ProfilePicture != nil {
		avatar, err := edit.s3.UploadToS3MultipartFileHeader(ProfilePicture, "userAvatar/"+user+".jpg")
		if err != nil {
			ctx.Redirect(http.StatusSeeOther, "/")
			return
		}
		if err := edit.update.UpdateBio(edit.userBioData, avatar); err != nil {
			ctx.JSON(http.StatusInternalServerError, false)
			return
		}
		ctx.Redirect(http.StatusSeeOther, "/"+user)
		return
	}
	if err := edit.update.UpdateBio(edit.userBioData, ""); err != nil {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	//  then go to /profile/:username
	ctx.Redirect(http.StatusSeeOther, "/"+user)

}

func (edit UserEdit) UserAddAcheivements(ctx *gin.Context) {
	user := ctx.GetString("user")
	if res := edit.check.CheckUser(user); !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if err := ctx.ShouldBind(&edit.userAddAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !"})
		return
	}
	edit.userAddAchievements.UserName = user
	if err := validate.Struct(edit.userAddAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !"})
		return
	}
	val := edit.achievement.GetNewAchievementName(user, edit.userAddAchievements.Content)
	if val == "" {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	location, err := edit.s3.UploadToS3MultipartFileHeader(edit.userAddAchievements.Data, "userAchievements/"+user+"_"+edit.userAddAchievements.Content+"_"+val+".jpg")
	if err != nil {
		edit.transaction.RollBackTransaction()
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	res := edit.update.InsertAchievements(val, location)
	if !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/"+user)

}

func (edit UserEdit) UserDelAcheivements(ctx *gin.Context) {
	user := ctx.GetString("user")
	if res := edit.check.CheckUser(user); !res {
		ctx.JSON(http.StatusSeeOther, gin.H{"status": false, "message": "Invalid token claims", "result": nil})
		return
	}
	if err := ctx.BindJSON(&edit.userDelAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !", "result": nil})
		return
	}
	edit.userDelAchievements.UserName = user
	if err := validate.Struct(edit.userDelAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !", "result": nil})
		return
	}
	res := edit.update.DeleteAchievement(edit.userDelAchievements.Data)
	if !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/"+user)

}

func (edit UserEdit) UserSocialEdit(ctx *gin.Context) {
	user := ctx.GetString("user")
	if res := edit.check.CheckUser(user); !res {
		ctx.JSON(http.StatusSeeOther, gin.H{"status": false, "message": "Invalid token claims", "result": nil})
		return
	}
	if err := ctx.BindJSON(&edit.userSocial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !", "result": nil})
		return
	}
	edit.userSocial.UserName = user
	if err := validate.Struct(edit.userSocial); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !", "result": nil})
		return
	}
	res := edit.update.UserSocialUpdate(edit.userSocial)
	if res != "true" {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/"+user)
}

func (edit UserEdit) UpdateNotification(ctx *gin.Context) {
	user := ctx.GetString("user")
	action := ctx.Param("action")
	if err := ctx.BindJSON(&edit.userNotification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !", "result": nil})
		return
	}
	edit.userNotification.Action = action
	if err := validate.Struct(edit.userNotification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid !", "result": nil})
		return
	}
	if status := edit.check.CheckUserHasClan(user); status {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "user is already in a clan"})
		return
	}
	res := edit.help.UpdateNotification(edit.userNotification.Id, edit.userNotification.Action)
	if !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/"+user)
}

func (edit UserEdit) SendTeamJoinRequest(ctx *gin.Context) {
	team := ctx.Param("team")
	user := ctx.GetString("user")
	if res := edit.check.CheckUserHasClan(user); res {
		ctx.JSON(http.StatusBadRequest, gin.H{"result": false, "message": "User is already in a clan"})
		return
	}
	if res := edit.update.InsertTeamNotification(user, team, "join clan"); !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request successfully completed"})
}

func (edit UserEdit) UserChangePassword(ctx *gin.Context) {

}
