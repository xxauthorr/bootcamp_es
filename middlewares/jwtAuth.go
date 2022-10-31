package middlewares

import (
	"bootcamp_es/services/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mwares struct {
	jwt  jwt.Jwt
	Token  *string
}

var X string

func (mw *Mwares) Authneticate(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("token")
	if clientToken == "" {
		ctx.Redirect(http.StatusFound, "/home")
		ctx.Abort()
		return
	}
	claims, err := mw.jwt.ValidateToken(clientToken)
	if err != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}
	X = claims.User
	ctx.Set("username", claims.User)
	ctx.Next()
}
