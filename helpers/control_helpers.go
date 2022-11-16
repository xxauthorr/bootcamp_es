package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
	"bootcamp_es/services/jwt"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Help struct {
	get        database.Get
	check      database.Check
	tokenCheck jwt.Jwt
}

func (h Help) GetPhone(username string) string {
	res := h.check.CheckUser(username)
	if res {
		return ""
	}

	phone := h.get.GetPhoneNumber(username)
	return phone
}

func (h Help) GetUsername(phone string) string {
	res := h.check.CheckPhoneNumber(phone)
	if !res {
		return ""
	}
	username := h.get.GetUsername(phone)
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
			if err == "signature is invalid" || err == "token expired" {
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
		if err == "signature is invalid" || err == "token expired" {
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

func (h Help) GetToken(user string) models.Token {
	var data models.Token
	token, expiresAt, refreshToken, err := h.tokenCheck.GenerateToken(user)
	if err != nil {
		fmt.Println("error at generating token:", err.Error())
	}
	data.AccessToken = token
	data.RefreshToken = refreshToken
	data.ExpiresAt = expiresAt
	return data
}

func (h Help) GetHomeData() models.HomeData {

	data := h.get.TopEntities()
	return data
	// should return the banners and the youtube results
}
