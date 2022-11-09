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

type Team struct {
	teamRegister models.TeamReg
	// teamDb 		database.Team
	checkDb database.Check
	team    helpers.TeamHelper
}
type EditTeam struct {
	edit                database.TeamProfileUpdate
	teamAddAchievements models.TeamAchievementsAdd
	bucket              amazons3.S3
	transaction         database.DBoperation
}

func (t Team) TeamProfle(ctx *gin.Context) {
	// teamname := ctx.Param("teamname")
	// if teamname == "" {
	// 	ctx.JSON(http.StatusNotFound, nil)
	// 	return
	// }
	// if res := t.checkDb.CheckTeam(teamname); !res {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "Team not found!"})
	// 	return
	// }
	// teamData := t.teamDb.FetchTeamData(teamname)

}

func (t Team) RegisterTeam(ctx *gin.Context) {
	user := middlewares.TokenUser
	if err := ctx.BindJSON(&t.teamRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.teamRegister.Username = user
	if err := validate.Struct(t.teamRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "Validation Error", "error": err.Error()})
		return
	}
	if status := t.checkDb.CheckTeam(t.teamRegister.TeamName); status {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "Team name already taken!"})
		return
	}
	// function check whether the user is already in a team, if not the team is registered and the notification to players are send
	msg, err := t.team.TeamScanAndInsert(t.teamRegister, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if msg != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": msg})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "team is registered"})
}

func (c EditTeam) TeamAddAchievements(ctx *gin.Context) {
	if err := ctx.ShouldBind(&c.teamAddAchievements); err != nil {
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	c.teamAddAchievements.TeamName = middlewares.Team
	if err := validate.Struct(c.teamAddAchievements); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	val := c.edit.GetAchievmentsName(c.teamAddAchievements)
	if val == "" {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	location, err := c.bucket.UploadToS3(c.teamAddAchievements.Data, c.teamAddAchievements.TeamName+"_"+c.teamAddAchievements.Content+"_"+val+".jpg")
	if err != nil {
		c.transaction.RollBackTransaction()
		fmt.Println("s3 error")
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	res := c.edit.InsertTeamAchievements(location,val)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Data insertion error"})
		return
	}

}
