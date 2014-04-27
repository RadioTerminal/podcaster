package routes

import (
	"../models"
	"bytes"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"text/template"
)

const (
	ContentType = "Content-Type"
	PlainType   = "text/plain"
	RSSType     = "application/rss+xml"
)

type Feed struct {
	Podcast *models.Group
	Media   *[]models.Media
}

var xml_feed_t = `<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:media="http://search.yahoo.com/mrss/" version="2.0">
	<channel>
		<title>{{.Podcast.Name}}</title>
		<link>http://podcasti.radioterminal.si/#/podcast/{{.Podcast.Slug}}</link>
		<language>{{.Podcast.Language}}</language>
		<description>{{.Podcast.Description}}</description>
		<itunes:subtitle>{{.Podcast.Description}}</itunes:subtitle>
		<itunes:summary>{{.Podcast.Text | html}}</itunes:summary>
		<itunes:category text="Music"></itunes:category>
		<itunes:author>{{.Podcast.Author}}</itunes:author>
		<itunes:owner>
			<itunes:email>{{.Podcast.Email}}</itunes:email>
			<itunes:name>{{.Podcast.Author}}</itunes:name>
		</itunes:owner>
		<itunes:explicit>{{.Podcast.Explicit}}</itunes:explicit>
		<image>
			<url>{{.Podcast.PictureUrl}}</url>
			<title>{{.Podcast.Name}}</title>
			<link>http://podcasti.radioterminal.si/#/podcast/{{.Podcast.Slug}}</link>
		</image>
		<itunes:image href="{{.Podcast.PictureUrl}}" />
		<copyright>{{.Podcast.Author}}</copyright>{{$slug := .Podcast.Slug}}{{$author := .Podcast.Author}}{{$email := .Podcast.Email}}{{with .Media}}{{range .}}
		<item>
			<title>{{.Name}}</title>
			<guid isPermaLink="false">http://podcasti.radioterminal.si/#/podcast/{{$slug}}/{{.Id}}</guid>
			<pubDate>{{.CreatedAt}}</pubDate>
			<author>{{$email}} ({{$author}})</author>
			<itunes:author>{{$author}}</itunes:author>
			<description>{{.Text | html}}</description>
			<itunes:summary>{{.Text | html}}</itunes:summary>
			<itunes:subtitle>{{.Text | short}}...</itunes:subtitle>
			<itunes:keywords>{{.Tags}}</itunes:keywords>
			<media:thumb url="{{.CoverUrl}}" />
			<enclosure url="{{.Url}}" length="{{.Duration}}" type="audio/mpeg" />
			<media:content lang="en" medium="audio" url="{{.Url}}" expression="full" fileSize="{{.Duration}}" type="audio/mpeg" isDefault="true" />
			<image>
				<url>{{.CoverUrl}}</url>
				<title>{{.Name}}</title>
			</image>
			<itunes:image href="{{.CoverUrl}}" />
			<link>http://podcasti.radioterminal.si/#/podcast/{{$slug}}/{{.Id}}</link>
		</item>{{end}}{{end}}
	</channel>
</rss>`

func ShortenString(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	return s[:250]
}

func FeedForGroupGet(db gorm.DB, params martini.Params, res http.ResponseWriter) (int, string) {
	podcast := models.Group{}
	var media = []models.Media{}
	if err := db.Where(&models.Group{Slug: params["slug"]}).First(&podcast).Error; err != nil {
		res.Header().Set(ContentType, PlainType)
		return http.StatusNotFound, "Feed not found"
	}
	db.Model(&podcast).Order("created_at desc").Limit(15).Related(&media)

	tpl, err := template.New("xml_feed").Funcs(template.FuncMap{"short": ShortenString}).Parse(xml_feed_t)
	if err != nil {
		res.Header().Set(ContentType, PlainType)
		return http.StatusInternalServerError, "Generating XML failed"
	}
	data := Feed{&podcast, &media}
	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, data)
	log.Println(err)
	if err != nil {
		res.Header().Set(ContentType, PlainType)
		return http.StatusInternalServerError, "Generating XML failed"
	}
	res.Header().Set(ContentType, RSSType)
	return http.StatusOK, buf.String()
}
