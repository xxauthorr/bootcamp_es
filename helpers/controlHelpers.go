package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/services/jwt"
	"fmt"

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

func (h Help) Authorize(ctx *gin.Context) bool {
	clientToken := ctx.Request.Header.Get("token")
	if clientToken == "" {
		clientRefreshToken := ctx.Request.Header.Get("refresh_token")
		if clientRefreshToken == "" {
			return false
		}
		claims, err := h.tokenCheck.ValidateRefreshToken(clientRefreshToken)
		if err != "" {
			fmt.Println(err)
			if err == "signature is invalid" || err == "token is expired" {
				return false
			}
			// should log the error
			return false
		}
		ctx.Set("user", claims.User)
		return true
	}
	claims, err := h.tokenCheck.ValidateToken(clientToken)
	if err != "" {
		fmt.Println(err, "error")
		if err == "signature is invalid" || err == "token is expired" {
			return false
		}
		// should log the error
		return false
	}
	ctx.Set("user", claims.User)
	return true
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

// func (h Help) Search(data models.Search){
// 	if data.Entity == "user"{

// 	}
// }
