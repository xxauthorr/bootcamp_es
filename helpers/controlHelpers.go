package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/services/jwt"

	"github.com/gin-gonic/gin"
)

type Help struct {
	getUserDetails database.Get
	check          database.Check
	tokenCheck     jwt.Jwt
}

// func (h *Help) DelChar(s []rune, index int) []rune {
// 	return append(s[0:index], s[index+1:]...)
// }

func (h Help) GetPhone(username string) string {
	res := h.check.CheckUser(username)
	if res {
		return ""
	}

	phone := h.getUserDetails.GetPhoneNumber(username)
	return phone
}

func (h *Help) GetUsername(phone string) string {
	res := h.check.CheckPhoneNumber(phone)
	if !res {
		return ""
	}
	username := h.getUserDetails.GetUsername(phone)
	return username
}

func (h *Help) Authneticate(ctx *gin.Context) (bool, string) {
	clientToken := ctx.Request.Header.Get("token")
	if clientToken == "" {
		return false, ""
	}
	claims, err := h.tokenCheck.ValidateToken(clientToken)
	if err != "" {
		return false, ""
	}

	ctx.Set("username", claims.User)
	return true, claims.User
}
