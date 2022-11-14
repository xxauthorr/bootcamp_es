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
	db          database.Tournament
	transaction database.DBoperation
	helper      helpers.Tournament
}

func (c Tour) TournamentRegistration(ctx *gin.Context) {
	user := ctx.GetString("user")
	c.register.Game = ctx.Request.PostFormValue("game")
	c.register.Tournament_name = ctx.Request.PostFormValue("name")
	c.register.Prize_pool = ctx.Request.PostFormValue("prizepool")
	c.register.No_of_slots = ctx.Request.PostFormValue("slots")
	c.register.Registration_ends = ctx.Request.PostFormValue("reg_ends")
	c.register.T_start = ctx.Request.PostFormValue("start")
	c.register.T_end = ctx.Request.PostFormValue("end")
	c.register.Registration_link = ctx.Request.PostFormValue("reg_link")
	c.register.Live_stream = ctx.Request.PostFormValue("live")
	c.register.Discord = ctx.Request.PostFormValue("discord")
	// if err := ctx.BindXML(&c.register); err != nil {
	// 	fmt.Println(err.Error(),"error here")
	// 	ctx.Redirect(http.StatusSeeOther, "/"+user)
	// 	return
	// }
	c.register.User = user
	fmt.Println(c.register.Live_stream, c.register.Prize_pool)
	if err := validate.Struct(c.register); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid entry !", "result": nil})
		return
	}
	res, id := c.db.RegisterTournament(c.register)
	if !res {
		c.transaction.RollBackTransaction()
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	if res := c.helper.RegistrartionTournamentFiles(ctx, id); !res {
		ctx.Redirect(http.StatusSeeOther, "/")
		return
	}
	ctx.Redirect(http.StatusOK, "/"+user)
}
