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

type Team struct {
	teamRegister models.TeamReg
	teamDb       database.Team
	checkDb      database.Check
	team         helpers.TeamHelper
	getHelp      helpers.Help
	get          database.Get
	result       models.AuthResult
}
type EditTeam struct {
	edit                database.TeamProfileUpdate
	check               database.Check
	teamAddAchievements models.TeamAchievementsAdd
	teamDelAchievements models.TeamAchievementsDel
	teamNotification    models.Notification
	teamBioEdit         models.TeamBioEdit
	bucket              amazons3.S3
	transaction         database.DBoperation
}

func (t Team) TeamProfile(ctx *gin.Context) {
	teamname := ctx.Param("teamname")
	if res := t.checkDb.CheckTeam(teamname); !res {
		//team not found
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Team not found!", "result": nil})
		return
	}
	res := t.getHelp.Authorize(ctx)
	user := ctx.GetString("user")
	leader := t.get.GetTeamLeader(teamname)

	teamData := t.teamDb.FetchTeamData(teamname)
	teamData.Visit = true
	t.result.Data = teamData

	if leader == user {
		teamData = t.teamDb.FetchTeamNotification(teamData)
		teamData.Visit = false
		t.result.Data = teamData
		t.result.User = user
		t.result.Authorization = t.getHelp.GetToken(user)
		// show team data for the owner
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "successfully compelted", "result": t.result})
		return
	}
	if !res {
		ctx.JSON(http.StatusOK, gin.H{"status:": true, "message": "successfully completed", "result": t.result})
		return
	}
	t.result.Authorization = t.getHelp.GetToken(user)
	t.result.User = user
	ctx.JSON(http.StatusOK, gin.H{"status:": true, "message": "successfully completed", "result": t.result})
}

func (t Team) RegisterTeam(ctx *gin.Context) {
	user := ctx.GetString("user")
	if err := ctx.BindJSON(&t.teamRegister); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	t.teamRegister.Username = user
	if err := validate.Struct(t.teamRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid entry !"})
		return
	}
	if status := t.checkDb.CheckUserHasClan(user); status {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "User is already in a team !"})
		return
	}
	if status := t.checkDb.CheckTeam(t.teamRegister.TeamName); status {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Team name already taken!"})
		return
	}
	//Team is registered and the notification to players are send
	err := t.team.TeamScanAndInsert(t.teamRegister, user)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+t.teamRegister.TeamName)
}

func (c EditTeam) TeamEditBio(ctx *gin.Context) {
	team := ctx.GetString("team")
	if err := ctx.ShouldBind(&c.teamBioEdit); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	c.teamBioEdit.TeamName = team
	if err := validate.Struct(c.teamBioEdit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid entry !"})
		return
	}
	_, picture, _ := ctx.Request.FormFile("avatar")
	if picture != nil {
		location, err := c.bucket.UploadToS3MultipartFileHeader(picture, "teamAvatar/"+c.teamBioEdit.TeamName+".jpg")
		if err != nil {
			//should log error
			ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
			return
		}
		if res := c.edit.UpdateBio(c.teamBioEdit, location); !res {
			ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
			return
		}
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	if res := c.edit.UpdateBio(c.teamBioEdit, ""); !res {
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
}

func (c EditTeam) TeamAddAchievements(ctx *gin.Context) {
	content := ctx.Param("content")
	team := ctx.GetString("team")
	if err := ctx.ShouldBind(&c.teamAddAchievements); err != nil {
		fmt.Println(err.Error())
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	c.teamAddAchievements.TeamName = team
	if content == "tournament" {
		c.teamAddAchievements.Content = "TOURNAMENT"
	}
	if content == "scrims" {
		c.teamAddAchievements.Content = "SCRIMS"
	} else {
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	if err := validate.Struct(c.teamAddAchievements); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid entry !"})
		return
	}
	val := c.edit.GetAchievmentsName(c.teamAddAchievements)
	if val == "" {
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	location, err := c.bucket.UploadToS3MultipartFileHeader(c.teamAddAchievements.Data, "/teamAchievements/"+c.teamAddAchievements.TeamName+"_"+c.teamAddAchievements.Content+"_"+val+".jpg")
	if err != nil {
		c.transaction.RollBackTransaction()
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	res := c.edit.InsertTeamAchievements(location, val)
	if !res {
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Achievements successfully inserted"})
}

func (c EditTeam) TeamDelAchievements(ctx *gin.Context) {
	team := ctx.GetString("team")
	if err := ctx.BindJSON(&c.teamDelAchievements); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	c.teamDelAchievements.TeamName = team
	if err := validate.Struct(c.teamDelAchievements); err != nil {
		fmt.Println("error in validate")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "error in validation", "result": nil})
		return
	}
	if res := c.edit.DeleteTeamAchievements(c.teamDelAchievements.Data); !res {
		fmt.Println("error in delete")
		ctx.JSON(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
}

func (c EditTeam) UpdateTeamNotification(ctx *gin.Context) {
	team := ctx.GetString("team")
	if err := ctx.BindJSON(&c.teamNotification); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	c.teamNotification.Action = ctx.Param("action")
	if err := validate.Struct(c.teamNotification); err != nil {
		fmt.Println("error in validate")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "error in validation", "result": nil})
		return
	}
	if res := c.edit.UpdateTeamNotification(c.teamNotification, team); !res {
		fmt.Println("error in update Team Notification")
		ctx.Redirect(http.StatusSeeOther, "/teamprofile/"+team)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request completed successfully"})
}

func (c EditTeam) SendTeamJoinRequeset(ctx *gin.Context) {
	user := ctx.Param("user")
	team := ctx.GetString("team")
	if res := c.check.CheckTeam(team); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Team doesn't exist"})
		return
	}
	if err := c.edit.InsertTeamNotification(user, team, "Member"); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request succesfully completed"})
}
