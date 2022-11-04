package middlewares

import (
	"bootcamp_es/services/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mwares struct {
	jwt jwt.Jwt
}

var TokenUser string

func (mw *Mwares) AuthneticateToken(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("token")
	clientRefreshToken := ctx.Request.Header.Get("referesh_token")
	if clientToken == "" {
		ctx.Redirect(http.StatusFound, "/home")
		ctx.Abort()
		return
	}

	claims, err := mw.jwt.ValidateAccessToken(clientToken)
	if err != "" {
		if err == "signature is invalid" {
			ctx.Redirect(http.StatusPermanentRedirect, "/home")
			return
		}
		if err == "token is expired" {
			claims, err := mw.jwt.ValidateRefreshToken(clientRefreshToken)
			if err != "" {
				if err == "signature is invalid" {
					ctx.Redirect(http.StatusPermanentRedirect, "/home")
					return
				}
				if err == "token is expired" {
					ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "Refresh token expired, you must login"})
					return
				}
				ctx.JSON(http.StatusOK, gin.H{"error": err})
				return
			}
			TokenUser = claims.User
			ctx.Set("username", claims.User)
			ctx.Next()
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}
	TokenUser = claims.User
	ctx.Set("username", claims.User)
	ctx.Next()
}
