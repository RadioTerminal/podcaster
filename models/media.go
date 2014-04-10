package models

import (
	"time"
)

type Media struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Text string `json:"text"`

	Tags string `json:"tags"`

	User string `json:"user"`

	Url string `json:"url"`

	PictureUrl string `json:"img"`

	CreatedAt time.Time `json:"created"`

	UpdatedAt time.Time `json:"updated"`
}
