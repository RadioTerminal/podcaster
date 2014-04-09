package routes

import (
	"../models"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/gorilla/feeds"
	"net/http"
	"strconv"
	"time"
)

func GetGroupIndex(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var media []models.Group
	_, err := db.Select(&media, "select * from group order by id")
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(groupToIface(media)...))
}

func GetGroup(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Group{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Group)
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func AddGroup(entity models.Group, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	err := db.Insert(&entity)
	if err != nil {
		checkErr(err, "insert failed")
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/podcaster/group/%d", entity.Id))
	return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateGroup(entity models.Group, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Group{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	oldEntity := obj.(*models.Group)

	entity.Id = oldEntity.Id
	_, err = db.Update(&entity)
	if err != nil {
		checkErr(err, "update failed")
		return http.StatusConflict, ""
	}
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteGroup(db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Group{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Group)
	_, err = db.Delete(entity)
	if err != nil {
		checkErr(err, "delete failed")
		return http.StatusConflict, ""
	}
	return http.StatusNoContent, ""
}

func GetFeed(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	slug := parms["slug"]
	var group models.Group
	err := db.SelectOne(&group, "select * from group where slug=?", slug)
	if err != nil {
		// Invalid slug, or does not exist
		checkErr(err, "missing group")
		return http.StatusNotFound, ""
	}
	//entity := obj.(*models.Media)
	now := time.Now()
	feed := &feeds.Feed{
		Title:       group.Name,
		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
		Created:     now,
	}
	var media []models.Media
	err_media := db.SelectOne(&media, "select * from media where group=?", slug)
	if err_media != nil {
		checkErr(err_media, "No media")
		return http.StatusNotFound, ""
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
		checkErr(err_parsing, "generating failed")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, rss
}

func groupToIface(v []models.Group) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}
