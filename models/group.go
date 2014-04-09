package models

type Group struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Text string `json:"text"`

	Author string `json:"author"`

	PictureUrl string `json:"picture"`
}
