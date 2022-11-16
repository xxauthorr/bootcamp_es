package controllers

import (
	"bootcamp_es/database"
	"bootcamp_es/helpers"
	"bootcamp_es/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tour struct {
	register    models.Tournament_registration_data
	check       database.Check
	db          database.Tournament
	transaction database.DBoperation
	helper      helpers.Tournament
}

func (c Tour) TournamentRegistration(ctx *gin.Context) {
	user := ctx.GetString("user")
	if err := ctx.ShouldBind(&c.register); err != nil {
		fmt.Println("error in bson bind :", err.Error())
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	c.register.User = user
	if err := validate.Struct(c.register); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid entry !", "result": nil})
		return
	}
	if res := c.check.CheckTournament(c.register.Tournament_name); res {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Tournament with this name alreaddy registered"})
		return
	}
	res, id := c.db.RegisterTournament(c.register)
	if !res {
		c.transaction.RollBackTransaction()
		fmt.Println("err1")
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if res := c.helper.RegisterTournamentFiles(ctx, id); !res {
		fmt.Println("err2")
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request completed successfully"})
}

