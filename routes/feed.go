package routes

import (
	"../models"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/yosssi/gold"
	"net/http"
)

const (
	ContentType = "Content-Type"
	PlainType   = "text/plain"
	RSSType     = "text/plain"
)

func FeedForGroupGet(db gorm.DB, params martini.Params, res http.ResponseWriter) (int, string) {
	podcast := models.Group{}
	media := models.Media{}
	if err := db.Where(&models.Group{Slug: params["slug"]}).First(&podcast).Error; err != nil {
		res.Header().Set(ContentType, PlainType)
		return http.StatusNotFound, "Feed not found"
	}
	db.Model(&podcast).Order("created_at desc").Limit(15).Related(&media)

	if err_parsing != nil {
		res.Header().Set(ContentType, PlainType)
		return http.StatusInternalServerError, "Generating XML failed"
	}
	res.Header().Set(ContentType, RSSType)
	return http.StatusOK, rss
}
