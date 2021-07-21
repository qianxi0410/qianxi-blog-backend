package model

import "time"

type Post struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Pre         int       `json:"pre"`
	Next        int       `json:"next"`
	Url         string    `json:"url"`
	Path        string    `json:"path"`
	Tags        string    `json:"tags"`
}

type PostWrap struct {
	Post
	NextTitle string `json:"next_title"`
	PreTitle  string `json:"pre_title"`
}
