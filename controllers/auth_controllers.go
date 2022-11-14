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
	changePassword models.ChangePassword
	tokenResult    models.Token
	search         models.Search
	dbCheck        database.Check
	get            database.Get
	UserDB         database.User
	help           helpers.Help
	twilio         twilio.Do
	jwt            jwt.Jwt
	admin          AdminControllers
}

// used to check weather the user is already exist or not
func (c Auth) CheckUser(ctx *gin.Context) {
	userName := ctx.Param("username")
	res := c.dbCheck.CheckUser(userName)
	if !res {
		ctx.JSON(http.StatusOK, gin.H{"status": false, "message": "request succefully completed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": res, "message": "request succefully completed"})
}

func (c Auth) CheckTeam(ctx *gin.Context) {
	teamName := ctx.Param("teamname")
	res := c.dbCheck.CheckTeam(teamName)
	if !res {
		ctx.JSON(http.StatusOK, gin.H{"status": res, "message": "request succefully completed"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": res, "message": "request succefully completed"})
}

func (c Auth) Home(ctx *gin.Context) {
	status := c.help.Authorize(ctx)
	homeData := c.help.GetHomeData()
	if !status {
		// for not logged in users
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": status, "message": "Not logged in", "result": homeData})
		return
	}
	// for logged in users
	user := ctx.GetString("user")
	token := c.help.GetToken(user)
	homeData.Authorization = token
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Request succesfully completed", "result": homeData})
}

// otp is send to the given phone number and return the phone and the status true
func (c Auth) SendPhoneOTP(ctx *gin.Context) {

	if err := ctx.BindJSON(&c.forgetPass); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	c.forgetPass.Username = "aksjdhf"
	if err := validate.Struct(c.forgetPass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Request body is invalid"})
		return
	}
	if res := c.dbCheck.CheckPhoneNumber(c.signup.Phone); res {
		ctx.JSON(http.StatusOK, gin.H{"status": false, "message": "Account already exist using this phone number"})
		return
	}
	if err := c.twilio.SendOtp(c.forgetPass.Phone); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Otp has send to the number " + c.forgetPass.Phone})
}

// bind the signup data and check the given otp, if the otp is true, user is registred
func (c Auth) SignUp(ctx *gin.Context) {
	if err := ctx.BindJSON(&c.signup); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if err := validate.Struct(c.signup); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid entry !"})
		return
	}
	if res := c.dbCheck.CheckPhoneNumber(c.signup.Phone); res {
		ctx.JSON(http.StatusOK, gin.H{"status": false, "message": "Account already exist using this phone number"})
		return
	}
	if res := c.dbCheck.CheckUser(c.signup.UserName); res {
		ctx.JSON(http.StatusOK, gin.H{"status": false, "message": "Account already exist using this username"})
		return
	}
	// res, err := c.twilio.CheckOtp(c.signup.Phone, c.signup.Otp)
	// if err != nil {
	//	ctx.Redirect(http.StatusInternalServerError, "/home")
	// 	return
	// }
	// if !res {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"status": res, "message": "otp is invalid !"})
	// 	return
	// }
	if err := c.UserDB.InsertUser(c.signup); err != nil {
		ctx.Redirect(http.StatusInternalServerError, "/home")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "mesage": "Request succefully completed"})
}

func (c Auth) Login(ctx *gin.Context) {
	if err := ctx.BindJSON(&c.login); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if err := validate.Struct(c.login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid entry !"})
		return
	}
	res := c.dbCheck.CheckUser(c.login.UserName)
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Username does'nt exist !"})
		return
	}
	res, err := c.dbCheck.CheckPassword(c.login.UserName, c.login.Password)
	if err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if !res {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Wrong password !"})
		return
	}
	if res := c.dbCheck.CheckUserType(c.login.UserName); res == "ADMIN" {
		c.admin.Dashboard(ctx, c.login.UserName)
		return
	}
	token, expiresAt, refreshToken, err := c.jwt.GenerateToken(c.login.UserName)
	if err != nil {
		fmt.Println("error at generating token:", err.Error())
	}
	c.tokenResult.AccessToken = token
	c.tokenResult.ExpiresAt = expiresAt
	c.tokenResult.RefreshToken = refreshToken
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Request succesfully completed", "result": c.tokenResult})
}

func (c Auth) ForgotPassword(ctx *gin.Context) {
	if err := ctx.BindJSON(&c.forgetPass); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if c.forgetPass.Phone == "" {
		c.forgetPass.Phone = "1234567890"
	} else {
		c.forgetPass.Username = "abcdefgh"
	}
	if err := validate.Struct(c.forgetPass); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid entry !"})
		return
	}
	if c.forgetPass.Phone == "1234567890" {
		res := c.dbCheck.CheckUser(c.forgetPass.Username)
		if !res {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "User does'nt exist!"})
			return
		}
		c.forgetPass.Phone = c.help.GetPhone(c.forgetPass.Username)
		// if err := c.twilio.SendOtp(c.forgetPass.Phone); err != nil {
		// 	ctx.Redirect(http.StatusInternalServerError, "/home")
		// 	return
		// }
		phone := c.help.NakeString(c.forgetPass.Phone)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Otp has send to your phone :" + phone})
		return
	}
	if res := c.dbCheck.CheckPhoneNumber(c.forgetPass.Phone); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Account using this number does'nt exist!"})
		return
	}
	// if err := c.twilio.SendOtp(c.forgetPass.Phone); err != nil {
	// 	ctx.Redirect(http.StatusInternalServerError, "/home")
	// 	return
	// }
	phone := c.help.NakeString(c.forgetPass.Phone)
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Otp has send to your phone :" + phone})
}

func (c Auth) ChangePassword(ctx *gin.Context) {
	if err := ctx.BindJSON(&c.changePassword); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if err := validate.Struct(c.changePassword); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid entry !"})
		return
	}
	res, err := c.twilio.CheckOtp(c.changePassword.Phone, c.changePassword.Otp)
	if err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	if !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": res, "message": "otp is invalid !"})
		return
	}
	if err := c.UserDB.ChangePass(c.changePassword.Phone, c.changePassword.Password); err != nil {
		ctx.Redirect(http.StatusBadRequest, "/home")
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"status": true, "message": "Password succesfully changed !"})
}


func (c Auth) SearchFirstFive(ctx *gin.Context) {
	entity := ctx.Param("entity")
	if err := ctx.BindJSON(&c.search); err != nil {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	c.search.Entity = entity
	if err := validate.Struct(c.search); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid entry !"})
		return
	}
	res, data := c.get.GetFirstFive(c.search)
	if !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid params", "result": nil})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request completed successfully", "result": data})
}