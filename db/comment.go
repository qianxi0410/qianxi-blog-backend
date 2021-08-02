package db

import (
	"github.com/qianxi/blog-backend/model"
)

type CommentDB struct{}

func (c CommentDB) Save(comment model.Comment) uint {
	db.Create(&comment)
	return comment.ID
}

func (c CommentDB) Get(postId uint) []model.Comment {
	var result []model.Comment

	db.Table("comments").Where("post_id", postId).Find(&result)
	return result
}

func (c CommentDB) Delete(id uint) {
	db.Delete(&model.Comment{}, id)
}
