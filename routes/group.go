package routes

import (
	"../models"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
	//"github.com/gorilla/feeds"
	//"time"
)

func GroupsIndex(db gorm.DB, r render.Render) {
	projects := []models.Group{}
	db.Find(&projects)
	r.JSON(http.StatusOK, projects)
}

// func GetFeed(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
// 	slug := parms["slug"]
// 	var group models.Group
// 	err := db.SelectOne(&group, "select * from group where slug=?", slug)
// 	if err != nil {
// 		// Invalid slug, or does not exist
// 		checkErr(err, "missing group")
// 		return http.StatusNotFound, ""
// 	}
// 	//entity := obj.(*models.Media)
// 	now := time.Now()
// 	feed := &feeds.Feed{
// 		Title:       group.Name,
// 		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
// 		Description: "discussion about tech, footie, photos",
// 		Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
// 		Created:     now,
// 	}
// 	var media []models.Media
// 	err_media := db.SelectOne(&media, "select * from media where group=?", slug)
// 	if err_media != nil {
// 		checkErr(err_media, "No media")
// 		return http.StatusNotFound, ""
// 	}
// 	feed.Items = []*feeds.Item{
// 		&feeds.Item{
// 			Title:       "Limiting Concurrency in Go",
// 			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
// 			Description: "A discussion on controlled parallelism in golang",
// 			Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
// 			Created:     now,
// 		},
// 		&feeds.Item{
// 			Title:       "Logic-less Template Redux",
// 			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
// 			Description: "More thoughts on logicless templates",
// 			Created:     now,
// 		},
// 		&feeds.Item{
// 			Title:       "Idiomatic Code Reuse in Go",
// 			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
// 			Description: "How to use interfaces <em>effectively</em>",
// 			Created:     now,
// 		},
// 	}
// 	rss, err_parsing := feed.ToRss()
// 	if err_parsing != nil {
// 		checkErr(err_parsing, "generating failed")
// 		return http.StatusInternalServerError, ""
// 	}
// 	return http.StatusOK, rss
// }
