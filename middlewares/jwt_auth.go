package middlewares

import (
	"bootcamp_es/services/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mwares struct {
	jwt         jwt.Jwt
	userChecker UserCheckers
}

// var UserChecker UserCheckers

func (mw Mwares) AuthneticateToken(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("token")
	var count int
	for i := range clientToken {
		if clientToken[i] == '.' {
			count = count + 1
		}
	}
	if count != 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Invalid token"})
		ctx.Abort()
		return
	}
	if clientToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Token not passed"})
		ctx.Abort()
		return
	}
	claims, err := mw.jwt.ValidateToken(clientToken)
	if err != "" {
		if err == "signature is invalid" || err == "token expired" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		ctx.Abort()
		return
	}
	ctx.Set("user", claims.User)
	mw.userChecker.CheckUserBlocked(ctx)
	ctx.Next()
}
