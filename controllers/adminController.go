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
	userData   models.UserProfileData
	entities   models.Entities
	search     models.Search
	check      database.Check
}

func (c AdminControllers) Dashboard(ctx *gin.Context, user string) {
	c.entities = c.helper.GetEntitiesCount()

	ctx.JSON(http.StatusOK, gin.H{"type": "admin", "Entities count": c.entities})
}

func (c AdminControllers) Search(ctx *gin.Context) {
	if err := ctx.BindJSON(&c.search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}
	if err := validate.Struct(c.search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}
	if c.search.Entity == "user" {
		if res := c.check.CheckUser(c.search.Value); !res {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "User Doesn't Exist!"})
			return
		}
		c.userData = c.userHelper.FetchUserData(c.search.Value)
		ctx.JSON(http.StatusOK, c.userData)
		return
	}
	if c.search.Entity == "team"{
		if res := c.check.CheckTeam(c.search.Value); !res {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "msg": "Team Doesn't Exist!"})
			return
		}
		
	}
	c.helper.AdminSearch(c.search)
}
