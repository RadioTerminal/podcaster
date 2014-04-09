package models

type Media struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Text string `json:"text"`

	Tags string `json:"tags"`

	User string `json:"user"`

	Url string `json:"url"`

	Group int64 `json:"group"`

	PictureUrl string `json:"img"`
}
