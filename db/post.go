package db

import (
	"fmt"

	"github.com/qianxi/blog-backend/model"
)

type PostDB struct{}

func (p PostDB) Get(id uint) model.Post {
	var result model.Post
	db.First(&result, id)

	return result
}

func (p PostDB) GetPostByPageQuery(page, size int) []model.Post {
	var result []model.Post
	offset := (page - 1) * size
	db.Table("posts").Offset(offset).Limit(size).Order("created_at DESC").Find(&result)

	return result
}

func (p PostDB) GetPostByPageAndTagQuery(page, size int, tag string) []model.Post {
	var result []model.Post
	offset := (page - 1) * size
	db.Table("posts").Where("tags LIKE ?", fmt.Sprintf("%%%s%%", tag)).Offset(offset).Limit(size).Order("created_at DESC").Find(&result)

	return result
}

func (p PostDB) Count() int64 {
	var result int64
	db.Table("posts").Count(&result)

	return result
}

func (p PostDB) CountWithTag(tag string) int64 {
	var result int64
	db.Table("posts").Where("tags LIKE ?", fmt.Sprintf("%%%s%%", tag)).Count(&result)

	return result
}
