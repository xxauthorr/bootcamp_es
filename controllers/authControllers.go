package controllers

import (
	// "bootcamp_es/models"
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"
	"bootcamp_es/services/twilio"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Controller struct {
	ForOtp models.ForOtp
	Check  database.Check
	UserDB database.User
	help   helpers.Help
	twilio twilio.Do
}

func (do *Controller) SendPhoneOTP(ctx *gin.Context) {
	// var phone models.ForOtp
	if err := ctx.BindJSON(&do.ForOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Struct(do.ForOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// To remove the zeroth index from the given string
	*do.ForOtp.Number = string(do.help.DelChar([]rune(*do.ForOtp.Number), 0))
	// Checking whether the number already exits or not
	if err := do.Check.CheckPhoneNumber(*do.ForOtp.Number); err != nil {
		fmt.Println(err.Error())
		if err.Error() == "true" {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Given number already exist !"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calling twilio service to send otp
	if err := do.twilio.SendOtp(*do.ForOtp.Number); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Take me to enter otp"})
}

func (do *Controller) CheckPhoneOtp(ctx *gin.Context) {

	number := *do.ForOtp.Number
	if err := ctx.BindJSON(&do.ForOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	do.ForOtp.Number = &number
	fmt.Println(*do.ForOtp.Number)
	fmt.Println(*do.ForOtp.Otp)
	if err := validate.Struct(do.ForOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	status, err := do.twilio.CheckOtp(*do.ForOtp.Number, *do.ForOtp.Otp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !status {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Entered wrong do !"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Take me to the signup page", "phone": do.ForOtp.Number})

}

// used to check weather the user is already exist or not
func CheckUser(ctx *gin.Context) {
	var username string
	if err := ctx.BindJSON(&username); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(username)
	ctx.JSON(http.StatusOK, gin.H{"username": username})

}

func (do *Controller) Signup(ctx *gin.Context) {
	var user models.Signup
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := do.Check.CheckUser(*user.UserName); err != nil {
		if err.Error() == "Exist" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Username already taken"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// do jwtt !!!!
	if err := do.UserDB.RegisterUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Take me to login page"})
}

func (do Controller) Login(ctx *gin.Context) {

}
