package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminControllers struct {
	helper         helpers.AdminHelper
	help           helpers.Help
	search         models.Search
	admin          database.Admin
	result         models.AdminResult
	updateUserData models.UpdateUserType
	check          database.Check
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
	user := ctx.GetString("user")
	if err := ctx.BindJSON(&c.search); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		return
	}
	if err := validate.Struct(c.search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid body"})
		return
	}
	res, data := c.admin.GetSerachData(c.search)
	temp := fmt.Sprint(reflect.TypeOf(data))
	if temp == "models.Search" {
		if res != "" {
			if res == "invalid" {
				ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid request"})
				return
			}
			if res == "no data" {
				ctx.JSON(http.StatusOK, gin.H{"status": false, "message": c.search.Entity + " with this name does'nt exist"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
			return
		}

	}
	c.result.Authorize = c.help.GetToken(user)
	c.result.Admin = user
	c.result.Data = data
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request succesfully completed", "result": c.result})
}

func (c AdminControllers) ListUsers(ctx *gin.Context) {
	user := ctx.GetString("user")
	val := ctx.Param("page")
	page, _ := strconv.ParseInt(val, 6, 12)
	data := c.admin.GetUsersList(page)
	if fmt.Sprint(reflect.TypeOf(data)) == "models.AdminUserData" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data, "user": user})
}

func (c AdminControllers) ListTeam(ctx *gin.Context) {
	user := ctx.GetString("user")
	val := ctx.Param("page")
	page, _ := strconv.ParseInt(val, 6, 12)
	data := c.admin.GetTeamList(page)
	if fmt.Sprint(reflect.TypeOf(data)) == "models.AdminTeamData" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "user": user, "data": data})
}

func (c AdminControllers) ListTournament(ctx *gin.Context) {
	user := ctx.GetString("user")
	val := ctx.Param("page")
	page, _ := strconv.ParseInt(val, 6, 12)
	data := c.admin.GetTournamentList(page)
	if fmt.Sprint(reflect.TypeOf(data)) == "models.AdminTournamentData" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "user": user, "data": data})
}

func (c AdminControllers) UpdateUserType(ctx *gin.Context) {
	action := ctx.Param("action")
	if err := ctx.BindJSON(&c.updateUserData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid body"})
		return
	}
	c.updateUserData.Action = action
	if err := validate.Struct(c.updateUserData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid body"})
		return
	}
	user := ctx.GetString("user")
	if res := c.check.CheckUser(c.updateUserData.User); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "User doesn't exit"})
		return
	}
	if res := c.admin.MakeAdmin(c.updateUserData.Action, c.updateUserData.User); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request successfully completed", "user": user})
}

func (c AdminControllers) UpdateBlock(ctx *gin.Context) {
	action := ctx.Param("action")
	if err := ctx.BindJSON(&c.updateUserData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid body"})
		return
	}
	c.updateUserData.Action = action
	if err := validate.Struct(c.updateUserData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid body"})
		return
	}
	user := ctx.GetString("user")
	if res := c.check.CheckUser(c.updateUserData.User); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "User doesn't exit"})
		return
	}
	if res := c.admin.Block(c.updateUserData.Action, c.updateUserData.User); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request successfully completed", "user": user})
}
