package routes

import (
	"../models"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
)

func MediaIndex(db gorm.DB, r render.Render) {
	media := []models.Media{}
	db.Find(&media)
	r.JSON(http.StatusOK, media)
}

func MediaPlay(db gorm.DB, r render.Render, params martini.Params) {
	media := models.Media{}
	id, _ := strconv.Atoi(params["id"])
	if err := db.First(&media, id).Error; err != nil {
		r.Error(http.StatusNotFound)
		return
	}
	db.Model(&media).Update(&models.Media{Played: media.Played + 1})
	r.Redirect(media.Url)
}

func MediaGet(db gorm.DB, r render.Render, params martini.Params) {
	media := models.Media{}
	if err := db.Where("slug = ?", params["slug"]).First(&media).Error; err != nil {
		r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Media not found"})
		return
	}
	r.JSON(http.StatusOK, media)
}

func MediaNew(r render.Render) {
	media := new(models.Media)
	r.JSON(http.StatusOK, media)
}

func MediaCreate(db gorm.DB, r render.Render, media models.Media) {
	if err := db.Save(&media).Error; err != nil {
		r.JSON(http.StatusConflict, map[string]interface{}{"error": "Media conflict"})
		return
	}
	r.JSON(http.StatusCreated, media)
}

func MediaUpdate(db gorm.DB, r render.Render, params martini.Params, updatedMedia models.Media) {
	var media models.Media
	if err := db.First(&media, params["id"]).Error; err != nil {
		r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Media not found"})
		return
	}
	db.Model(&media).Update(&updatedMedia)
	r.JSON(http.StatusOK, media)
}

func MediaDelete(db gorm.DB, r render.Render, params martini.Params) {
	var media models.Media
	if err := db.First(&media, params["id"]).Error; err != nil {
		r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Media not found"})
		return
	}
	db.Delete(media)
	r.JSON(http.StatusNoContent, nil)
}
