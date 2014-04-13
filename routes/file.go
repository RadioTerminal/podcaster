package routes

import (
	"../models"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
)

func FileUpload(db gorm.DB, r render.Render) {

	r.JSON(http.StatusCreated, media)
}
