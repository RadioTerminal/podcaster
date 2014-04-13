package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
)

func FileUpload(r render.Render) {

	r.JSON(http.StatusCreated, "ok")
}
