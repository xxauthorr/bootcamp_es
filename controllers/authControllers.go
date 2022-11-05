package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"

	// bycrypt "bootcamp_es/services/byCrypt"
	"bootcamp_es/services/jwt"

	"bootcamp_es/services/twilio"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Auth struct {
	signup         models.SignupForm
	login          models.LoginForm
	forgetPass     models.ForgetPassword
	forOtp         models.ForOtp
	changePassword models.ChangePassword
	dbCheck        database.Check
	UserDB         database.User
	help           helpers.Help
	twilio         twilio.Do
	token          jwt.Jwt
}

// used to check weather the user is already exist or not
func (do Auth) CheckUser(ctx *gin.Context) {
	username := ctx.Param("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "false", "msg": "No username has given !"})
		return
	}
	res := do.dbCheck.CheckUser(username)
	if !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": res})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": res})
}

func (do Auth) Home(ctx *gin.Context) {
	status, user := do.help.Authneticate(ctx)
	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": status, "msg": "You must login"})
		return
	}
	if user == "" {
		ctx.JSON(200, gin.H{"status": false})
		return
	}
	// defer func() {
	// 	if e := recover(); e != nil {

	// 		fmt.Println(e, "jdhgsd")
	// 	}
	// }()
	ctx.JSON(200, gin.H{"status": true, "user": user})
}

// otp is send to the given phone number and return the phone and the status true
func (do Auth) SendPhoneOTP(ctx *gin.Context) {

	if err := ctx.BindJSON(&do.forgetPass); err != nil {
		ctx.JSON(http.StatusBadRequest, "/home")
		return
	}
	do.forgetPass.Username = "aksjdhf"
	if err := validate.Struct(do.forgetPass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "phone string is empty"})
		return
	}
	if res := do.dbCheck.CheckPhoneNumber(do.signup.Phone); res {
		ctx.JSON(http.StatusOK, gin.H{"status": false, "msg": "Account already exist using this phone number"})
		return
	}
	if err := do.twilio.SendOtp(do.forgetPass.Phone); err != nil {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "true", "phone": do.forgetPass.Phone})
}

// bind the signup data and check the given otp, if the otp is true, user is registred
func (do Auth) SignUp(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.signup); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if err := validate.Struct(do.signup); err != nil {
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	if res := do.dbCheck.CheckPhoneNumber(do.signup.Phone); res {
		ctx.JSON(http.StatusOK, gin.H{"status": false, "msg": "Account already exist using this phone number"})
		return
	}
	if res := do.dbCheck.CheckUser(do.signup.UserName); res {
		ctx.JSON(http.StatusOK, gin.H{"status": false, "msg": "Account already exist using this username"})
		return
	}
	// res, err := do.twilio.CheckOtp(do.signup.Phone, do.signup.Otp)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, false)
	// 	return
	// }
	// if !res {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": res, "msg": "otp is invalid !"})
	// 	return
	// }
	if err := do.UserDB.InsertUser(do.signup); err != nil {
		ctx.Redirect(http.StatusInternalServerError, "/home")
		return
	}
	ctx.JSON(http.StatusOK, true)
}

func (do Auth) Login(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.login); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if err := validate.Struct(do.login); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := do.dbCheck.CheckUser(do.login.UserName)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Username does'nt exist !"})
		return
	}
	res, err := do.dbCheck.CheckPassword(do.login.UserName, do.login.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong Password"})
		return
	}

	token, _, err := do.token.GenerateToken(do.login.UserName)
	if err != nil {
		fmt.Println("error at generating token:", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "next": "take me to home", "token": token})
}

func (do Auth) ForgotPassword(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.forgetPass); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if do.forgetPass.Phone == "" {
		do.forgetPass.Phone = "1234567890"
	} else {
		do.forgetPass.Username = "abcdefgh"
	}
	if err := validate.Struct(do.forgetPass); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if do.forgetPass.Phone == "1234567890" {
		res := do.dbCheck.CheckUser(do.forgetPass.Username)
		if !res {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "user does'nt exist!"})
			return
		}
		do.forgetPass.Phone = do.help.GetPhone(do.forgetPass.Username)
		if err := do.twilio.SendOtp(do.forgetPass.Phone); err != nil {
			ctx.Redirect(http.StatusInternalServerError, "/home")
			return
		}
		phone := do.help.NakeString(do.forgetPass.Phone)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "msg": "otp has send to your phone :" + phone})
		return
	}
	if res := do.dbCheck.CheckPhoneNumber(do.forgetPass.Phone); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "account using this number does'nt exist!"})
		return
	}
	if err := do.twilio.SendOtp(do.forgetPass.Phone); err != nil {
		ctx.Redirect(http.StatusInternalServerError, "/home")
		return
	}
	phone := do.help.NakeString(do.forgetPass.Phone)
	ctx.JSON(http.StatusOK, gin.H{"status": true, "msg": "otp has send to your phone :" + phone})
}

func (do Auth) VerifyForgetOtp(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.forOtp); err != nil {
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	if err := validate.Struct(do.forOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	res, err := do.twilio.CheckOtp(do.forOtp.Number, do.forOtp.Otp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, false)
		return
	}
	if !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": res, "msg": "otp is invalid !"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": res, "msg": "otp successfully confirmed"})
}

func (do Auth) ChangePassword(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.changePassword); err != nil {
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	if err := validate.Struct(do.changePassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if err := do.UserDB.ChangePass(do.changePassword.Phone, do.changePassword.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, true)
}
