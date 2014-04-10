package routes

import (
	"../models"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
)

func GroupsIndex(db gorm.DB, r render.Render) {
	group := []models.Group{}
	db.Find(&group)
	r.JSON(http.StatusOK, group)
}

func GroupGet(db gorm.DB, r render.Render, params martini.Params) {
	podcast := models.Group{}
	if err := db.Where("slug = ?", params["slug"]).First(&podcast).Error; err != nil {
		r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Group not found"})
		return
	}
	r.JSON(http.StatusOK, podcast)
}

func MediaForGroupGet(db gorm.DB, r render.Render, params martini.Params) {
	podcast := models.Group{}
	media := models.Media{}
	if err := db.Where("slug = ?", params["slug"]).First(&podcast).Error; err != nil {
		r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Group not found"})
		return
	}
	files := db.Model(&podcast).Related(&media)
	r.JSON(http.StatusOK, files)
}

func GroupNew(r render.Render) {
	group := new(models.Group)
	r.JSON(http.StatusOK, group)
}

func GroupCreate(db gorm.DB, r render.Render, group models.Group) {
	if err := db.Save(&group).Error; err != nil {
		r.JSON(http.StatusConflict, map[string]interface{}{"error": "Group conflict"})
		return
	}
	r.JSON(http.StatusCreated, group)
}

func GroupUpdate(db gorm.DB, r render.Render, params martini.Params, updatedGroup models.Group) {
	var group models.Group
	if err := db.First(&group, params["id"]).Error; err != nil {
		r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Group not found"})
		return
	}
	db.Model(&group).Update(&updatedGroup)
	r.JSON(http.StatusOK, group)
}

func GroupDelete(db gorm.DB, r render.Render, params martini.Params) {
	var group models.Group
	if err := db.First(&group, params["id"]).Error; err != nil {
		r.JSON(http.StatusNotFound, map[string]interface{}{"error": "Group not found"})
		return
	}
	db.Delete(group)
	r.JSON(http.StatusNoContent, nil)
}
