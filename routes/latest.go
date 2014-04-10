package routes

import (
	"../models"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
)

func PopularIndex(db gorm.DB, r render.Render) {
	media := []models.Media{}
	db.Order("played desc").Limit(15).Find(&media)
	r.JSON(http.StatusOK, media)
}

func LatestIndex(db gorm.DB, r render.Render) {
	media := []models.Media{}
	db.Order("created_at desc").Limit(15).Find(&media)
	r.JSON(http.StatusOK, media)
}
