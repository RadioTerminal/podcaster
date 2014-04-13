package main

import (
	"./models"
	"./routes"
	"./utils"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/pilu/fresh/runner/runnerutils"
	"github.com/superlogical/analytics"
	"log"
	"net/http"
	"os"
	"time"
)

// The one and only martini instance.
var m *martini.Martini

func runnerMiddleware(w http.ResponseWriter, r *http.Request) {
	if runnerutils.HasErrors() {
		runnerutils.RenderError(w)
	}
}

var PublicKey []byte
var PrivateKey []byte

func init() {
	envFileName := martini.Env + ".env"
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Fatalf("Error loading: %s", envFileName)
	}
	utils.GenKeyPairIfNone(os.Getenv("PRIVATE_KEY"), os.Getenv("PUBLIC_KEY"))
	PrivateKey = utils.GetKey(os.Getenv("PRIVATE_KEY"))
	PublicKey = utils.GetKey(os.Getenv("PUBLIC_KEY"))
	claims := map[string]interface{}{
		"user_id": 1,
		"role":    "admin",
		"exp":     time.Now().UTC().Add(time.Hour * 6).Unix(),
		"iat":     time.Now().UTC().Unix(),
	}
	auth, _ := utils.GenerateAuthToken(claims, PrivateKey)
	fmt.Println(auth)

	m = martini.New()
	// Setup middleware
	if os.Getenv("DEV_RUNNER") == "1" {
		m.Use(runnerMiddleware)
	}
	m.Use(analytics.Google("UA-xxxxxxxx-1"))
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

		r.Get(`/media`, routes.MediaIndex)
		r.Get(`/media/play/:id`, routes.MediaPlay)
		r.Post(`/media`, utils.LoginRequired(PublicKey), binding.Bind(models.Media{}), routes.MediaCreate)
		r.Get("/media/new", utils.LoginRequired(PublicKey), routes.MediaNew)
		r.Get(`/media/:slug`, routes.MediaGet)
		r.Put(`/media/:id`, utils.LoginRequired(PublicKey), binding.Bind(models.Media{}), routes.MediaUpdate)
		r.Delete(`/media/:id`, utils.LoginRequired(PublicKey), routes.MediaDelete)

		r.Get(`/groups`, routes.GroupsIndex)
		r.Post(`/groups`, utils.LoginRequired(PublicKey), binding.Bind(models.Group{}), routes.GroupCreate)
		r.Get("/groups/new", utils.LoginRequired(PublicKey), routes.GroupNew)
		r.Get(`/group/:slug`, routes.GroupGet)
		r.Get(`/group/:slug/media`, routes.MediaForGroupGet)
		r.Put(`/group/:id`, utils.LoginRequired(PublicKey), binding.Bind(models.Group{}), routes.GroupUpdate)
		r.Delete(`/group/:id`, utils.LoginRequired(PublicKey), routes.GroupDelete)
	})

	gr.Get(`/feed/:slug`, routes.FeedForGroupGet)
	// Inject database

	// Add the router action
	m.Action(gr.Handle)
}

func main() {
	m.Run()
}
