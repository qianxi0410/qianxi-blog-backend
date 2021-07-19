package db

import "github.com/qianxi/blog-backend/model"

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
