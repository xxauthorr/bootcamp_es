package controllers

import (
	"github.com/xxauthorr/bootcamp_es/database"
	"github.com/xxauthorr/bootcamp_es/helpers"
	"github.com/xxauthorr/bootcamp_es/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tour struct {
	register    models.Tournament_registration_data
	editTour    models.EditTournamentData
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

func (c Tour) EditTournamentData(ctx *gin.Context) {
	if err := ctx.ShouldBind(&c.editTour); err != nil {
		fmt.Println("reached")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid body"})
		return
	}
	c.editTour.Name = ctx.GetString("tournament")
	if err := validate.Struct(c.editTour); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid body"})
		return
	}
	res := make(chan bool)
	go c.helper.UpdateTournamentFiles(ctx, c.editTour.Name, res)
	if response := c.db.UpdateTournament(c.editTour); !response {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "request successfully completed"})
}
