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

func (p PostService) Get(id string) (model.PostWrap, error) {
	var result model.PostWrap
	numberId, err := strconv.Atoi(id)
	if err != nil {
		return model.PostWrap{}, err
	}
	if numberId < 0 {
		return model.PostWrap{}, errors.New("oops ! id can not under than zero")
	}

	result.Post = postDB.Get(uint(numberId))

	if result.Title == "" {
		return result, errors.New("oops ! post not found ")
	}

	f, err := ioutil.ReadFile(result.Path)
	if err != nil {
		return model.PostWrap{}, err
	}
	if result.Pre != -1 {
		result.PreTitle = postDB.Get(uint(result.Pre)).Title
	}
	if result.Next != -1 {
		result.NextTitle = postDB.Get(uint(result.Next)).Title
	}

	result.Path = string(f)
	comments := commentDB.Get(uint(numberId))

	if len(comments) == 0 {
		result.Comments = make([]model.Comment, 0)
	} else {
		result.Comments = append(result.Comments, comments...)
	}
	return result, nil
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
		return nil, errors.New("oops ! there must be something err with page query ")
	}
	return result, nil
}

func (p PostService) GetPostByPageAndTagQuery(page, size, tag string) ([]model.Post, error) {
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

	if len(tag) == 0 {
		return nil, errors.New("oops ! your tag is empty")
	}

	result := postDB.GetPostByPageAndTagQuery(numberPage, numberSize, tag)

	if len(result) == 0 {
		return nil, errors.New("oops ! there must be something err with page query ")
	}
	return result, nil
}

func (p PostService) Count() (int64, error) {
	result := postDB.Count()

	if result < 0 {
		return -1, errors.New("oops ! there must be something wrong with count ")
	}

	return result, nil
}

func (p PostService) CountWithTag(tag string) (int64, error) {
	if len(tag) == 0 {
		return -1, errors.New("oops ! your tag is empty ")
	}

	result := postDB.CountWithTag(tag)

	if result < 0 {
		return -1, errors.New("oops ! there must be something wrong with tag count ")
	}
	return result, nil
}
