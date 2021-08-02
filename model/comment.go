package model

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	Content   string    `json:"content"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	PostId    uint      `json:"post_id"`
}
