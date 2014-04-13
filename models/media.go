package models

import (
	"../utils"
	"errors"
	"github.com/extemporalgenome/slug"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/binding"
	"time"
)

type Media struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Text string `json:"text"`

	GroupId Group `sql:"not null" json:"group"`

	Tags string `json:"tags"`

	Author string `json:"author"`

	Played int64 `json:"play_count"`

	Duration float64 `json:"duration"`

	Waveform []string `json:"wave"`

	Url string `json:"url"`

	CoverUrl string `json:"cover_url"`

	CreatedAt time.Time `json:"created"`

	UpdatedAt time.Time `json:"updated"`
}

func (u Media) Validate(errors *binding.Errors) {
	// TODO: Check for duplicate name
	if len(u.Name) < 0 {
		errors.Fields["name"] = "Name is required"
	}
	if len(u.Text) < 0 {
		errors.Fields["text"] = "Text is required"
	}
	if len(u.Author) < 0 {
		errors.Fields["author"] = "Author is required"
	}
	if len(u.CoverUrl) < 0 {
		errors.Fields["cover_url"] = "Picture is required"
	}
}

func (u *Media) BeforeCreate(tx *gorm.DB) (err error) {
	var count int
	tx.Model(u).Where("name = ?", u.Name).Count(&count)
	if count > 0 {
		err = errors.New("Media exists!")
		return
	}
	u.Slug = slug.Slug(u.Name)
	data, duration := utils.GenerateSamplesAsString(u.Url, 4)
	if !(duration > 0) {
		err = errors.New("File duration is invalid!")
		return
	}
	u.Waveform = data
	u.Duration = duration
	return
}

func (u *Media) BeforeUpdate(tx *gorm.DB) (err error) {
	var count int
	tx.Model(u).Where("name = ?", u.Name).Count(&count)
	if count > 1 {
		err = errors.New("Conflicting Name!")
		return
	}
	u.Slug = slug.Slug(u.Name)
	return
}
