package service

import (
	"errors"
	"io/ioutil"
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
		return model.Post{}, errors.New("oops ! id can not under than zero")
	}

	post := postDB.Get(uint(numberId))

	if post.Title == "" {
		return post, errors.New("oops ! post not found ")
	}

	f, err := ioutil.ReadFile(post.Path)
	if err != nil {
		return model.Post{}, err
	}

	post.Path = string(f)
	return post, nil
}

func (p PostService) GetPostByPageQuery(page, size string) ([]model.Post, error) {
	numberPage, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	numberSize, err := strconv.Atoi(size)
	if err != nil {
		return nil, err
	}

	if numberPage < 0 || numberSize < 0 {
		return nil, errors.New("oops ! page and size can not under than zero")
	}

	result := postDB.GetPostByPageQuery(numberPage, numberSize)
	if len(result) == 0 {
		return nil, errors.New("oops ! there must be something err ")
	}
	return result, nil
}

func (p PostService) Count() (int64, error) {
	result := postDB.Count()

	if result < 0 {
		return -1, errors.New("oops ! there must be something wrong with count")
	}

	return result, nil
}
