package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/qianxi/blog-backend/db"
	"github.com/qianxi/blog-backend/model"
	rdx "github.com/qianxi/blog-backend/redis"
)

type PostService struct{}

var postDB db.PostDB

func (p PostService) Get(id string) (model.PostWrap, error) {
	var result model.PostWrap

	res, err := rdb.Get(ctx, rdx.Post(id)).Result()

	if err != redis.Nil {
		json.Unmarshal([]byte(res), &result)
		return result, nil
	}

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

	r, _ := json.Marshal(result)

	rdb.Set(ctx, rdx.Post(id), r, time.Minute*3)
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
	var result []model.Post
	res, err := rdb.Get(ctx, rdx.PostWithPageAndSize(numberPage, numberSize)).Result()

	if err != redis.Nil {
		json.Unmarshal([]byte(res), &result)
		return result, nil
	}

	result = postDB.GetPostByPageQuery(numberPage, numberSize)

	if len(result) == 0 {
		return nil, errors.New("oops ! there must be something err with page query ")
	}
	r, _ := json.Marshal(result)
	rdb.Set(ctx, rdx.PostWithPageAndSize(numberPage, numberSize), r, time.Minute*1)
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

	var result []model.Post
	res, err := rdb.Get(ctx, rdx.PostWithPageAndSizeAndTag(numberPage, numberSize, tag)).Result()

	if err != redis.Nil {
		json.Unmarshal([]byte(res), &result)
		return result, nil
	}

	result = postDB.GetPostByPageAndTagQuery(numberPage, numberSize, tag)

	if len(result) == 0 {
		return nil, errors.New("oops ! there must be something err with page query ")
	}
	r, _ := json.Marshal(result)

	rdb.Set(ctx, rdx.PostWithPageAndSizeAndTag(numberPage, numberSize, tag), r, time.Minute*1)

	return result, nil
}

func (p PostService) Count() (int64, error) {
	res, err := rdb.Get(ctx, rdx.PostCount()).Result()
	if err != redis.Nil {
		r, _ := strconv.Atoi(res)
		return int64(r), nil
	}
	result := postDB.Count()

	if result < 0 {
		return -1, errors.New("oops ! there must be something wrong with count ")
	}

	rdb.Set(ctx, rdx.PostCount(), result, time.Hour)

	return result, nil
}

func (p PostService) CountWithTag(tag string) (int64, error) {
	if len(tag) == 0 {
		return -1, errors.New("oops ! your tag is empty ")
	}

	res, err := rdb.Get(ctx, rdx.TagCount(tag)).Result()

	if err != redis.Nil {
		r, _ := strconv.Atoi(res)
		return int64(r), nil
	}

	result := postDB.CountWithTag(tag)

	if result < 0 {
		return -1, errors.New("oops ! there must be something wrong with tag count ")
	}
	rdb.Set(ctx, rdx.TagCount(tag), result, time.Hour)
	return result, nil
}
