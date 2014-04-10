package models

import (
	"github.com/codegangsta/martini-contrib/binding"
	"time"
)

type Group struct {
	Id int64 `json:"id"`

	Name string `json:"name" binding:"required"`

	Slug string `json:"slug"`

	Text string `json:"text" binding:"required"`

	Author string `json:"author" binding:"required"`

	PictureUrl string `json:"picture"`

	CreatedAt time.Time

	UpdatedAt time.Time
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
