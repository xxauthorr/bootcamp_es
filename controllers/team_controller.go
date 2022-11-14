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
}
type EditTeam struct {
	edit                database.TeamProfileUpdate
	teamAddAchievements models.TeamAchievementsAdd
	teamDelAchievements models.TeamAchievementsDel
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
	user := ctx.GetString("user")
	leader := t.get.GetTeamLeader(teamname)
	teamData := t.teamDb.FetchTeamData(teamname)
	teamData.Visit = true
	if leader == user {
		teamData = t.teamDb.FetchTeamNotification(teamData)
		token := t.getHelp.GetToken(user)
		teamData.Visit = false
		teamData.Token = token
		// show team data for the owner
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "successfully compelted", "result": teamData})
		return
	}
	// show data for the client
	ctx.JSON(http.StatusOK, gin.H{"status:": true, "message": "successfully completed", "result": teamData})
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
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Team is succesfully registered"})
}

func (c EditTeam) TeamEditBio(ctx *gin.Context) {
	if err := ctx.ShouldBind(&c.teamBioEdit); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/team")
		return
	}
	c.teamBioEdit.TeamName = ctx.GetString("team")
	if err := validate.Struct(c.teamBioEdit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid entry !"})
		return
	}
	location, err := c.bucket.UploadToS3MultipartFileHeader(c.teamBioEdit.Avatar, "teamAvatar/"+c.teamBioEdit.TeamName+".jpg")
	if err != nil {
		//should log error
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if res := c.edit.UpdateBio(c.teamBioEdit, location); !res {
		ctx.Redirect(http.StatusSeeOther, "/team")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/team")
}

func (c EditTeam) TeamAddAchievements(ctx *gin.Context) {
	content := ctx.Param("content")
	if err := ctx.ShouldBind(&c.teamAddAchievements); err != nil {
		fmt.Println(err.Error())
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	c.teamAddAchievements.TeamName = ctx.GetString("team")
	if content == "tournament" {
		c.teamAddAchievements.Content = "TOURNAMENT"
	}
	if content == "scrims" {
		c.teamAddAchievements.Content = "SCRIMS"
	} else {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if err := validate.Struct(c.teamAddAchievements); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid entry !"})
		return
	}
	val := c.edit.GetAchievmentsName(c.teamAddAchievements)
	if val == "" {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	location, err := c.bucket.UploadToS3MultipartFileHeader(c.teamAddAchievements.Data, "/teamAchievements/"+c.teamAddAchievements.TeamName+"_"+c.teamAddAchievements.Content+"_"+val+".jpg")
	if err != nil {
		c.transaction.RollBackTransaction()
		fmt.Println("s3 error")
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	res := c.edit.InsertTeamAchievements(location, val)
	if !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Achievements successfully inserted"})
}

func (c EditTeam) TeamDelAchievements(ctx *gin.Context) {
	fmt.Println("team del ")
	if err := ctx.BindJSON(&c.teamDelAchievements); err != nil {
		fmt.Println("error in binding")
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	c.teamDelAchievements.TeamName = ctx.GetString("team")
	fmt.Println(c.teamDelAchievements.TeamName)
	if err := validate.Struct(c.teamDelAchievements); err != nil {
		fmt.Println("error in validate")
		ctx.JSON(http.StatusSeeOther, "/team")
		return
	}
	if res := c.edit.DeleteTeamAchievements(c.teamDelAchievements.Data); !res {
		fmt.Println("error in delete")
		ctx.JSON(http.StatusSeeOther, "/team")
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/team")
}
