package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
	"bootcamp_es/services/jwt"

	// "errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Help struct {
	getUserDetails database.Get
	check          database.Check
	jwt            models.ForJwt
	team           database.Team
	tokenCheck     jwt.Jwt
}

// func (h *Help) DelChar(s []rune, index int) []rune {
// 	return append(s[0:index], s[index+1:]...)
// }

func (h Help) GetPhone(username string) (error, string) {
	err := h.check.CheckUser(username)
	if err != nil {
		if err.Error() != "Exist" {
			return err, ""
		}
	}
	if err == nil {
		return nil, ""
	}
	phone := h.getUserDetails.GetPhoneNumber(username)
	return nil, phone
}

func (h *Help) GetUsername(phone string) (error, string) {
	err := h.check.CheckPhoneNumber(phone)
	if err != nil {
		if err.Error() != "Exist" {
			return err, ""
		}
	}
	if err == nil {
		return nil, ""
	}
	username := h.getUserDetails.GetUsername(phone)
	return nil, username
}

func (h *Help) CheckJwtUser(username string) bool {
	jwtUser := *h.jwt.User
	if jwtUser == username {
		return true
	}
	fmt.Println(jwtUser, "function CheckJwtUser")
	return false
}

func (h *Help) TeamScan(team models.TeamReg,user string) (string, error) {
	//check weather the leader already have a team
	// leader := *h.jwt.User
	// fmt.Println(leader)
	// status, err := h.check.TeamLeaderCheck(leader)
	// if err != nil {
	// 	return "", errors.New("error in checking leader")
	// }
	// if !status {
	// 	return "User already in a team", nil
	// }
	for i := range *team.Players {
		player := *team.Players
		fmt.Println(player[i])
		if err := h.team.InsertTeamNotification(player[i], *team.TeamName, "Member"); err != nil {
			return "", err
		}
	}
	if err := h.team.InsertTeamNotification(*team.CoLeader, *team.TeamName, "Co-Leader"); err != nil {
		return "", err
	}
	if err := h.team.RegisterTeam(team, "leader"); err != nil {
		return "", err
	}
	return "", nil
}

func (h *Help) Authneticate(ctx *gin.Context) string {
	clientToken := ctx.Request.Header.Get("token")
	if clientToken == "" {
		return ""
	}
	claims, err := h.tokenCheck.ValidateToken(clientToken)
	if err != "" {
		return ""
	}

	ctx.Set("username", claims.User)
	return claims.User
}
