package model

import (
	"time"

	"gorm.io/gorm"
)

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.DeletedAt = time.Date(1970, 1, 1, 0, 0, 0, 0, &time.Location{})
	return
}
