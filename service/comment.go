package service

import (
	"errors"
	"strconv"

	"github.com/qianxi/blog-backend/db"
	"github.com/qianxi/blog-backend/model"
)

type CommentService struct{}

var commentDB db.CommentDB

func (c CommentService) Save(comment model.Comment) (uint, error) {
	if len(comment.Content) >= 255 {
		return 0, errors.New("oops ! your comment are too long ")
	}

	if comment.Avatar == "" || comment.Login == "" || comment.Name == "" {
		return 0, errors.New("oops ! you are missing some info about you ")
	}

	result := commentDB.Save(comment)
	return result, nil
}

func (c CommentService) Delete(id string) error {
	numId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	commentDB.Delete(uint(numId))
	return nil
}
