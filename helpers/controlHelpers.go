package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
	"bootcamp_es/services/jwt"

	"github.com/gin-gonic/gin"
)

type Help struct {
	getUserDetails database.Get
	check          database.Check
	tokenCheck     jwt.Jwt
}

func (h Help) GetPhone(username string) string {
	res := h.check.CheckUser(username)
	if res {
		return ""
	}

	phone := h.getUserDetails.GetPhoneNumber(username)
	return phone
}

func (h Help) GetUsername(phone string) string {
	res := h.check.CheckPhoneNumber(phone)
	if !res {
		return ""
	}
	username := h.getUserDetails.GetUsername(phone)
	return username
}

func (h Help) Authneticate(ctx *gin.Context) (bool, string) {
	clientToken := ctx.Request.Header.Get("token")
	if clientToken == "" {
		return false, ""
	}
	// do like authenticate token
	//ctx.Request.Header.Get("refresh_token")
	claims, err := h.tokenCheck.ValidateAccessToken(clientToken)
	if err != "" {
		return false, ""
	}

	ctx.Set("username", claims.User)
	return true, claims.User
}

func (h Help) NakeString(value string) string {
	var val string = ""
	count := len(value) - 5
	for i := range value {
		if i >= count+4 {
			val = val + string(value[i])
		} else if i >= count {
			val = val + "*"
		} else {
			val = val + string(value[i])
		}
	}
	return val
}

func (h Help) Search(data models.Search){
	if data.Entity == "user"{
		
	}
}