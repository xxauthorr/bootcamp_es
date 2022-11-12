package middlewares

import (
	"bootcamp_es/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminCheckers struct {
	check database.Check
}
type TeamCheckers struct {
	get database.Get
}

func (c *AdminCheckers) CheckUserType(ctx *gin.Context) {
	tokenUser := ctx.GetString("user")
	if res := c.check.CheckUserType(tokenUser); res != "ADMIN" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Only Admin can Access!!"})
		return
	}
	ctx.Next()
}

func (c *TeamCheckers) CheckLeaderTeam(ctx *gin.Context) {
	tokenUser := ctx.GetString("user")
	team := c.get.GetTeamFromLeader(tokenUser)
	if team == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "User Don't have a team"})
		return
	}
	ctx.Set("team", team)
	ctx.Next()
}
