package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/middlewares"
	"bootcamp_es/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Team struct {
	teamRegister models.TeamReg
	checkDb      database.Check
	team         helpers.TeamHelper
}

func (t Team) CheckTeamName(ctx *gin.Context) {
	teamName := ctx.Param("teamname")
	res := t.checkDb.CheckTeam(teamName)
	if res {
		ctx.JSON(http.StatusOK, gin.H{"status": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": res})
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
