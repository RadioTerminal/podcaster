package main

import (
	"./models"
	"./routes"
	"github.com/go-martini/martini"
	"github.com/joho/godotenv"
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
		r.Get(`/latest`, routes.GroupsIndex)

		// r.Get(`/media`, routes.GetMediaIndex)
		// r.Get(`/media/:id`, routes.GetMedia)
		// r.Post(`/media`, binding.Json(models.Media{}), routes.AddMedia)
		// r.Put(`/media/:id`, binding.Json(models.Media{}), routes.UpdateMedia)
		// r.Delete(`/media/:id`, routes.DeleteMedia)

		// r.Get(`/groups`, routes.GetGroupIndex)
		// r.Get(`/group/:id`, routes.GetGroup)
		// r.Post(`/groups`, binding.Json(models.Group{}), routes.AddGroup)
		// r.Put(`/group/:id`, binding.Json(models.Group{}), routes.UpdateGroup)
		// r.Delete(`/group/:id`, routes.DeleteGroup)
	})

	//gr.Get(`/feeds/:slug`, routes.GetFeed)
	// Inject database

	// Add the router action
	m.Action(gr.Handle)
}

func main() {
	m.Run()
}
