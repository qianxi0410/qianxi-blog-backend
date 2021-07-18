package db

import "github.com/qianxi/blog-backend/model"

type PostDB struct{}

func (p PostDB) Get(id uint) model.Post {
	var result model.Post
	db.First(&result, id)

	return result
}
