package models

import (
	"errors"
	"github.com/extemporalgenome/slug"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/binding"
	"time"
)

type Group struct {
	Id int64 `json:"id"`

	Name string `json:"name" binding:"required"`

	Slug string `json:"slug"`

	Text string `json:"text" binding:"required"`

	Author string `json:"author" binding:"required"`

	PictureUrl string `json:"picture"`

	CreatedAt time.Time `json:"created"`

	UpdatedAt time.Time `json:"updated"`
}

func (u Group) Validate(errors *binding.Errors) {
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
}

func (u *Group) BeforeCreate(tx *gorm.DB) (err error) {
	var count int
	tx.Model(u).Where("name = ?", u.Name).Count(&count)
	if count > 0 {
		err = errors.New("Conflicting Name!")
		return
	}
	u.Slug = slug.Slug(u.Name)
	return
}

func (u *Group) BeforeUpdate(tx *gorm.DB) (err error) {
	var count int
	tx.Model(u).Where("name = ?", u.Name).Count(&count)
	if count > 1 {
		err = errors.New("Conflicting Name!")
		return
	}
	u.Slug = slug.Slug(u.Name)
	return
}
