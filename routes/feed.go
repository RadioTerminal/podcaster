package routes

import (
	"../models"
	"github.com/go-martini/martini"
	"github.com/gorilla/feeds"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"net/http"
	"time"
)

func FeedForGroupGet(db gorm.DB, r render.Render, params martini.Params) {
	podcast := models.Group{}
	//media := models.Media{}
	if err := db.Where(&models.Group{Slug: params["slug"]}).First(&podcast).Error; err != nil {
		r.Data(http.StatusNotFound, []byte("Group not found"))
		return
	}
	//files := db.Model(&podcast).Related(&media)

	now := time.Now()
	feed := &feeds.Feed{
		Title:       podcast.Name,
		Link:        &feeds.Link{Href: podcast.Slug},
		Description: podcast.Text,
		Author:      &feeds.Author{podcast.Author, podcast.Author},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&feeds.Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}
	rss, err_parsing := feed.ToRss()
	if err_parsing != nil {
		r.Data(http.StatusInternalServerError, []byte("Generating XML failed"))
		return
	}
	r.Data(http.StatusOK, []byte(rss))
}
