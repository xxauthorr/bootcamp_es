package middlewares

import (
	"bootcamp_es/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Team string

type Checkers struct {
	check database.Check
}

func (c Checkers) CheckUserType(ctx *gin.Context) {
	if res := c.check.CheckUserType(TokenUser); res != "ADMIN" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "Only Admin can Access!!"})
		return
	}
	ctx.Next()
}

func (c Checkers) CheckLeaderTeam(ctx *gin.Context) {
	Team = c.check.GetTeamFromLeader(TokenUser)
	if Team == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "User Don't have a team"})
		return
	}
	ctx.Next()
}
