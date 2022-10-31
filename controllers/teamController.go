package controllers

import (
	"bootcamp_es/helpers"
	"bootcamp_es/middlewares"
	"bootcamp_es/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Team struct {
	teamRegister models.TeamReg
	help         helpers.Help
}

func (t *Team) CheckTeamName(ctx *gin.Context) {

}

func (t *Team) RegisterTeam(ctx *gin.Context) {
	token := middlewares.X
	if err := ctx.BindJSON(&t.teamRegister); err != nil {
		fmt.Println(t.teamRegister.Players)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(t.teamRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "false", "msg": "Validation Error"})
		return
	}
	fmt.Println("registerteam")
	// function check whether the user is already in a team, if not the team is registered and the notification to players are send
	fmt.Println("error HERE")
	msg, err := t.help.TeamScan(t.teamRegister, token)
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
