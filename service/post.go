package service

import (
	"errors"
	"strconv"

	"github.com/qianxi/blog-backend/db"
	"github.com/qianxi/blog-backend/model"
)

type PostService struct{}

var postDB db.PostDB

func (p PostService) Get(id string) (model.Post, error) {
	numberId, err := strconv.Atoi(id)
	if err != nil {
		return model.Post{}, err
	}
	if numberId < 0 {
		return model.Post{}, errors.New("id can not under 0")
	}

	post := postDB.Get(uint(numberId))

	if post.Title == "" {
		return post, errors.New("post not found")
	}

	return post, nil
}
