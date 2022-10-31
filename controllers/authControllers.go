package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"
	"bootcamp_es/services/jwt"

	// "bootcamp_es/services/twilio"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Auth struct {
	forOtp     models.ForOtp
	signup     models.SignupForm
	login      models.LoginForm
	forgetPass models.ForgetPassword
	changePass models.ChangePassword
	dbCheck    database.Check
	UserDB     database.User
	help       helpers.Help
	// twilio     twilio.Do
	token jwt.Jwt
}

// used to check weather the user is already exist or not
func (do Auth) CheckUser(ctx *gin.Context) {
	username := ctx.GetString("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "false", "msg": "No username has given !"})
		return
	}
	err := do.dbCheck.CheckUser(username)
	if err != nil {
		if err.Error() != "Exist" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "false", "error": err.Error()})
			return
		}
	}
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "false", "msg": "username does'nt exist !"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "true"})
}

func (do *Auth) Home(ctx *gin.Context) {
	user := do.help.Authneticate(ctx);
	if user  == "" {
		ctx.JSON(200, gin.H{"status": false})
		return
	}
	// defer func() {
	// 	if e := recover(); e != nil {
		
	// 		fmt.Println(e, "jdhgsd")
	// 	}
	// }()
	ctx.JSON(200, gin.H{"status": true, "user": user })
}

//	func send otp to the given phone number
//	also checks weather the number exist or not

func (do *Auth) SendPhoneOTP(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.forOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	temp := "000000"
	do.forOtp.Otp = &temp
	if err := validate.Struct(do.forOtp); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("ok")
	// Checking whether the number already exits or not
	if err := do.dbCheck.CheckPhoneNumber(*do.forOtp.Number); err != nil {
		if err.Error() == "Exist" {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Given number already exist !"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calling twilio service to send otp
	// if err := do.twilio.SendOtp(*do.forOtp.Number); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"msg": "Take me to enter otp"})
}

// checks weather the given number is valid or not

func (do *Auth) CheckPhoneOtp(ctx *gin.Context) {

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
	// status, err := do.twilio.CheckOtp(*do.forOtp.Number, *do.forOtp.Otp)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// if !status {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Entered wrong do !"})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{"msg": "Take me to the signup page"})

}

func (do *Auth) Signup(ctx *gin.Context) {
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

func (do *Auth) Login(ctx *gin.Context) {
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

	token, refreshToken, err := do.token.GenerateToken(*do.login.UserName)
	if err != nil {
		fmt.Println("error at generating token:", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Logged In", "token": token, "refresh_token": refreshToken})
}

func (do *Auth) ForgotPassword(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.forgetPass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	if do.forgetPass.Phone == nil && do.forgetPass.Username != nil {
		temp := "123456789"
		do.forgetPass.Phone = &temp
	}
	if do.forgetPass.Phone != nil && do.forgetPass.Username == nil {
		temp := "username"
		do.forgetPass.Username = &temp
	}
	if err := validate.Struct(do.forgetPass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	// if the user gives the username to change password
	if *do.forgetPass.Phone == "123456789" {
		err, phone := do.help.GetPhone(*do.forgetPass.Username)
		if phone == "" {
			if err == nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Username does'nt exist !"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		do.forgetPass.Phone = &phone
		// Calling twilio service to send otp
		// if err := do.twilio.SendOtp(*do.forgetPass.Phone); err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }
		ctx.JSON(http.StatusOK, gin.H{"msg": "otp has been send to your number"})
		return
	}
	// user gives the phone to change password
	err, username := do.help.GetUsername(*do.forgetPass.Phone)
	if username == "" {
		if err == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Account with this phone number does'nt exist !"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	do.forgetPass.Username = &username
	// if err := do.twilio.SendOtp(*do.forgetPass.Phone); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	fmt.Println(*do.forgetPass.Phone, " ")
	ctx.JSON(http.StatusOK, gin.H{"msg": "otp has been send to your number"})
}

func (do *Auth) VerifyPassOtp(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.forOtp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	do.forOtp.Number = do.forgetPass.Phone
	if err := validate.Struct(do.forOtp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	// status, err := do.twilio.CheckOtp(*do.forOtp.Number, *do.forOtp.Otp)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// if !status {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "OTP does'nt match !"})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{"msg": "enter new password"})
}

func (do *Auth) ChangePassword(ctx *gin.Context) {
	if err := ctx.BindJSON(&do.changePass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	do.changePass.Phone = do.forOtp.Number
	if err := validate.Struct(do.changePass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}
	if err := do.UserDB.ChangePass(*do.changePass.Phone, *do.changePass.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "go to login page "})
}
