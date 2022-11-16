package helpers

import (
	"bootcamp_es/database"
	"bootcamp_es/models"
	amazons3 "bootcamp_es/services/AmazonS3"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Tournament struct {
	bucket      amazons3.S3
	db          database.Tournament
	transaction database.DBoperation
}

func (h Tournament) RegisterTournamentFiles(ctx *gin.Context, id int64) bool {
	var data models.Tournament_registration_data
	data.Id = id
	val := fmt.Sprint(id)
	_, banner, _ := ctx.Request.FormFile("banner")
	_, prize_pool, _ := ctx.Request.FormFile("prize")
	_, road_map, _ := ctx.Request.FormFile("road_map")

	if banner != nil {
		banner, err := h.bucket.UploadToS3MultipartFileHeader(banner, "tournament/banner/"+val+".jpg")
		if err != nil {
			fmt.Println("reached")
			h.transaction.RollBackTransaction()
			fmt.Println(err.Error(), "error here ")
			return false
		}
		data.Banner = banner
	}
	if prize_pool != nil {
		prize_pool, err := h.bucket.UploadToS3MultipartFileHeader(prize_pool, "tournament/prizepool/"+val+".jpg")
		if err != nil {
			h.transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false
		}
		data.Prize_pool_banner = prize_pool
	}
	if road_map != nil {
		road_map, err := h.bucket.UploadToS3MultipartFileHeader(road_map, "tournament/roadmap/"+val+".jpg")
		if err != nil {
			h.transaction.RollBackTransaction()
			fmt.Println(err.Error())
			return false
		}
		data.Road_map = road_map
	}
	res := h.db.UpdateFile(data)
	if !res {
		h.transaction.RollBackTransaction()
		return res
	}
	return true
}
