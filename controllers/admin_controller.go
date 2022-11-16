package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminControllers struct {
	helper     helpers.AdminHelper
	userHelper helpers.UserHelper
	help       helpers.Help
	search     models.Search
	check      database.Check
}

func (c AdminControllers) AdminHome(ctx *gin.Context) {
	user := ctx.GetString("user")
	c.Dashboard(ctx, user)
}

func (c AdminControllers) Dashboard(ctx *gin.Context, user string) {
	entities := c.helper.GetEntitiesCount()
	entities.Authorization = c.help.GetToken(user)
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Request completed succesfully", "Result": entities})
}

func (c AdminControllers) Search(ctx *gin.Context) {
	if err := ctx.BindJSON(&c.search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request body", "result": nil})
		return
	}
	if err := validate.Struct(c.search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request body", "result": nil})
		return
	}
	if c.search.Entity == "user" {
		if res := c.check.CheckUser(c.search.Value); !res {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "User Doesn't Exist!", "result": nil})
			return
		}
		userData := c.userHelper.FetchProfileData(c.search.Value, false)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "successfully completed", "result": userData})
		return
	}
	if c.search.Entity == "team" {
		if res := c.check.CheckTeam(c.search.Value); !res {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "Team Doesn't Exist!", "result": nil})
			return
		}
		//return team profile data
	}
}
