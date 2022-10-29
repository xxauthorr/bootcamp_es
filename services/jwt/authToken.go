package jwt

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Jwt struct {
}

type SignedDetails struct {
	User               string
	jwt.StandardClaims //Registered claims
}

func (j Jwt) loadEnv() (string, error) {
	// loads env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		return "", err
	}
	var key = os.Getenv("SECRET_KEY")

	return key, nil
}

func (j Jwt) GenerateToken(userName string) (string, string, error) {

	SECRET_KEY, err := j.loadEnv()
	if err != nil {
		return "", "", err
	}

	claims := &SignedDetails{
		User: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refereshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err.Error())
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refereshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return token, refreshToken, nil
}

func (j Jwt) ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	SECRET_KEY, err := j.loadEnv()
	if err != nil {
		log.Fatal(err)
		return
	}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		msg = err.Error()
		return
	}
	return claims, msg
}

func UpdateTokens(signedToken string, signedRefreshToken string, userName string) {

}
