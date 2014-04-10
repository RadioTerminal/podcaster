package routes

import (
	"../models"
	"encoding/xml"
	"github.com/go-martini/martini"
	"github.com/gorilla/feeds"
	"github.com/jinzhu/gorm"
	"net/http"
)

const (
	ContentType = "Content-Type"
	PlainType   = "text/plain"
	RSSType     = "text/plain"
)

type rssFeedXml struct {
	XMLName xml.Name `xml:"rss"`
	XMLBase string   `xml:"xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:media="http://search.yahoo.com/mrss/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:geo="http://www.w3.org/2003/01/geo/wgs84_pos#" xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0"`
	Version string   `xml:"version,attr"`
	Channel *feeds.RssFeed
}

func FeedForGroupGet(db gorm.DB, params martini.Params, res http.ResponseWriter) (int, string) {
	podcast := models.Group{}
	//media := models.Media{}
	if err := db.Where(&models.Group{Slug: params["slug"]}).First(&podcast).Error; err != nil {
		res.Header().Set(ContentType, PlainType)
		return http.StatusNotFound, "Feed not found"
	}
	//files := db.Model(&podcast).Related(&media)
	fee := feeds.RssFeed{
		Title: podcast.Name,
		//Link:  &feeds.Link{Href: podcast.Slug},
		//Docs:           "http://www.rssboard.org/rss-specification",
		Description: podcast.Description,
		//ItunesSubtitle: podcast.Description,
		//Language:       podcast.Language,
		//Author:         &feeds.Author{podcast.Author, podcast.Author},
		//Created: now,
	}
	feed := rssFeedXml{
		Channel: &fee,
	}

	rss, err_parsing := toXML(feed)
	if err_parsing != nil {
		res.Header().Set(ContentType, PlainType)
		return http.StatusInternalServerError, "Generating XML failed"
	}
	res.Header().Set(ContentType, RSSType)
	return http.StatusOK, rss
}

func toXML(feed rssFeedXml) (string, error) {
	data, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		return "", err
	}
	// strip empty line from default xml header
	s := xml.Header[:len(xml.Header)-1] + string(data)
	return s, nil
}
