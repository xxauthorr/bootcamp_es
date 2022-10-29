package middlewares

import (
	"bootcamp_es/models"
	"bootcamp_es/services/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Mwares struct {
	jwt   jwt.Jwt
	login models.ForJwt
}

func (mw Mwares) Authneticate(ctx *gin.Context) {
	clientToken := ctx.Request.Header.Get("token")
	if clientToken == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No authorization header provided"})
		ctx.Abort()
		return
	}
	claims, err := mw.jwt.ValidateToken(clientToken)
	if err != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		ctx.Abort()
		return
	}
	*mw.login.User = claims.User
	ctx.Set("username", claims.User)
	ctx.Next()
}
