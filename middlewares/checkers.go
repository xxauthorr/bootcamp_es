package middlewares

import (
	"bootcamp_es/database"
	"fmt"
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
		ctx.Abort()
		return
	}
	ctx.Next()
}

// checks team exists by the leader user 
func (c *TeamCheckers) CheckLeaderTeam(ctx *gin.Context) {
	tokenUser := ctx.GetString("user")
	team := c.get.CheckTeamExist(tokenUser)
	if team == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "User Don't have a team"})
		ctx.Abort()
		return
	}
	ctx.Set("team", team)
	ctx.Next()
}

func (c *TeamCheckers) CheckTeamStrength(ctx *gin.Context) {
	team := ctx.GetString("team")
	strength := c.get.GetTeamStrength(team)
	if strength < 15 {
		need := 15 - strength
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": team + " needs " + fmt.Sprint(need) + " more members to do the operation", "result": nil})
		ctx.Abort()
		return
	}
	ctx.Next()
}
