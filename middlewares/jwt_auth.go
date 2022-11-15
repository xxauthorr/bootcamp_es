package middlewares

import (
	"bootcamp_es/database"
	"bootcamp_es/services/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mwares struct {
	jwt   jwt.Jwt
	check database.Check
}

func (mw Mwares) AuthneticateToken(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("token")
	if clientToken == "" {
		clientRefreshToken := ctx.Request.Header.Get("referesh_token")
		if clientRefreshToken == "" {
			ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"status": false, "Message": "User must login to do the operation"})
			ctx.Abort()
			return
		}
		claims, err := mw.jwt.ValidateToken(clientRefreshToken)
		if err != "" {
			if err == "signature is invalid" || err == "token expired" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid refresh token"})
				ctx.Abort()
				return
			}
			ctx.Redirect(http.StatusSeeOther, "/")
			ctx.Abort()
			return
		}
		ctx.Set("user", claims.User)
		ctx.Next()
		return
	}

	claims, err := mw.jwt.ValidateToken(clientToken)
	if err != "" {
		if err == "signature is invalid" || err == "token expired" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
			ctx.Abort()
			return
		}
		ctx.Redirect(http.StatusSeeOther, "/")
		ctx.Abort()
		return
	}
	ctx.Set("user", claims.User)
	ctx.Next()
}

func (mw Mwares) CheckUserType(ctx *gin.Context) {
	tokenUser := ctx.GetString("user")
	if res := mw.check.CheckUserType(tokenUser); res != "ADMIN" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Only Admin Have Access!!"})
		return
	}
	ctx.Next()
}
