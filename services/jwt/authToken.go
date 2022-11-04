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

func (j Jwt) loadEnv() (string, string, error) {
	// loads env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file loading error - ", err)
		return "", "", err
	}
	var key1 = os.Getenv("ACCESS_KEY")
	var key2 = os.Getenv("REFERSH_KEY")

	return key1, key2, nil
}

func (j Jwt) GenerateToken(userName string) (string, string, error) {

	ACCESS_KEY, REFRESH_KEY, err := j.loadEnv()
	if err != nil {
		return "", "", err
	}

	claims := &SignedDetails{
		User: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}

	refereshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(ACCESS_KEY))
	if err != nil {
		log.Panic(err.Error())
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refereshClaims).SignedString([]byte(REFRESH_KEY))
	if err != nil {
		log.Panic(err.Error())
		return "", "", err
	}
	return token, refreshToken, nil
}

func (j Jwt) ValidateAccessToken(AccessToken string) (claims *SignedDetails, msg string) {
	ACCESS_KEY, _, err := j.loadEnv()
	if err != nil {
		log.Fatal(err)
		return
	}
	token, err := jwt.ParseWithClaims(
		AccessToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(ACCESS_KEY), nil
		},
	)
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		// msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		// msg = fmt.Sprintf("token expired")
		msg = "token expired"
		return
	}
	if err != nil {
		if err.Error() == "signature is invalid" {
			msg = err.Error()
			return
		}
	}
	// 	fmt.Println("expired")
	// 	msg = "token expired"
	// 	return
	// }
	return claims, msg
}

func (j Jwt) ValidateRefreshToken(AccessToken string) (claims *SignedDetails, msg string) {
	_, REFERSH_KEY, err := j.loadEnv()
	if err != nil {
		log.Fatal(err)
		return
	}
	token, err := jwt.ParseWithClaims(
		AccessToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(REFERSH_KEY), nil
		},
	)
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		// msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		// msg = fmt.Sprintf("token expired")
		msg = "token expired"
		return
	}
	if err != nil {
		if err.Error() == "signature is invalid" {
			msg = err.Error()
			return
		}
	}
	return claims, msg
}

func UpdateTokens(signedToken string, signedRefreshToken string, userName string) {

}
