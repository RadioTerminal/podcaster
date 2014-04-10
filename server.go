package main

import (
	"./models"
	"./routes"
	"github.com/go-martini/martini"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/pilu/fresh/runner/runnerutils"
	"log"
	"net/http"
	"os"
)

// The one and only martini instance.
var m *martini.Martini

func runnerMiddleware(w http.ResponseWriter, r *http.Request) {
	if runnerutils.HasErrors() {
		runnerutils.RenderError(w)
	}
}

func init() {

	envFileName := martini.Env + ".env"
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Fatalf("Error loading: %s", envFileName)
	}

	m = martini.New()
	// Setup middleware
	if os.Getenv("DEV_RUNNER") == "1" {
		m.Use(runnerMiddleware)
	}
	m.Use(gzip.All())
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(render.Renderer())
	m.Use(martini.Static("public"))
	m.Map(models.Connect())
	// Setup routes
	gr := martini.NewRouter()
	gr.Group("/api", func(r martini.Router) {

		r.Get(`/latest`, routes.LatestIndex)
		r.Get(`/popular`, routes.PopularIndex)

		r.Get(`/media`, routes.GroupsIndex)
		r.Post(`/media`, binding.Bind(models.Group{}), routes.GroupCreate)
		r.Get("/media/new", routes.GroupNew)
		r.Get(`/media/:slug`, routes.GroupGet)
		r.Put(`/media/:id`, binding.Bind(models.Group{}), routes.GroupUpdate)
		r.Delete(`/media/:id`, routes.GroupDelete)

		r.Get(`/groups`, routes.GroupsIndex)
		r.Post(`/groups`, binding.Bind(models.Group{}), routes.GroupCreate)
		r.Get("/groups/new", routes.GroupNew)
		r.Get(`/group/:slug`, routes.GroupGet)
		r.Get(`/group/:slug/media`, routes.MediaForGroupGet)
		r.Put(`/group/:id`, binding.Bind(models.Group{}), routes.GroupUpdate)
		r.Delete(`/group/:id`, routes.GroupDelete)
	})

	gr.Get(`/feed/:slug`, routes.FeedForGroupGet)
	// Inject database

	// Add the router action
	m.Action(gr.Handle)
}

func main() {
	m.Run()
}
