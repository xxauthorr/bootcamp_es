package controllers

import (
	// "bootcamp_es/models"
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"
	"bootcamp_es/services/jwt"
	"bootcamp_es/services/twilio"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Controller struct {
	forOtp  models.ForOtp
	signup  models.SignupForm
	login   models.LoginForm
	dbCheck database.Check
	UserDB  database.User
	help    helpers.Help
	twilio  twilio.Do
	token   jwt.Jwt
}

func (do *Controller) SendPhoneOTP(ctx *gin.Context) {
	// var phone models.forOtp
	if err := ctx.BindJSON(&do.forOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(do.forOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Checking whether the number already exits or not
	if err := do.dbCheck.CheckPhoneNumber(string(do.help.DelChar([]rune(*do.forOtp.Number), 0))); err != nil {
		if err.Error() == "true" {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Given number already exist !"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calling twilio service to send otp
	if err := do.twilio.SendOtp(*do.forOtp.Number); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Take me to enter otp"})
}

func (do *Controller) CheckPhoneOtp(ctx *gin.Context) {

	number := *do.forOtp.Number
	if err := ctx.BindJSON(&do.forOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	do.forOtp.Number = &number
	if err := validate.Struct(do.forOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	status, err := do.twilio.CheckOtp(*do.forOtp.Number, *do.forOtp.Otp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !status {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Entered wrong do !"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Take me to the signup page"})

}

// used to check weather the user is already exist or not
func CheckUser(ctx *gin.Context) {
	var username string
	if err := ctx.BindJSON(&username); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"username": username})

}

func (do *Controller) Signup(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.signup); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	do.signup.Phone = do.forOtp.Number
	if err := validate.Struct(do.signup); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := do.dbCheck.CheckUser(*do.signup.UserName); err != nil {
		if err.Error() == "Exist" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Username already taken"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// do jwtt !!!!
	token, refresToken, err := do.token.GenerateToken(*do.signup.UserName)
	if err != nil {
		fmt.Println("error at generating token:", err.Error())
	}

	if err := do.UserDB.RegisterUser(do.signup); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "You are logged in", "token": token, "refresh_token": refresToken})
}

func (do Controller) Login(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.login); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := validate.Struct(do.login); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	// if do.forJwt.User != do.login.UserName {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid entry, Suspecious !"})
	// 	return
	// }
	err := do.dbCheck.CheckUser(*do.login.UserName)
	if err == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Username does'nt exist !"})
		return
	}
	if err.Error() != "Exist" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	res, err := do.dbCheck.CheckPassword(*do.login.UserName, *do.login.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong Password"})
		return
	}

	token, refresToken, err := do.token.GenerateToken(*do.login.UserName)
	if err != nil {
		fmt.Println("error at generating token:", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Logged In", "token": token, "refresh_token": refresToken})
}
