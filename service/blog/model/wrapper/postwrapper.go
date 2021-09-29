package wrapper

import "qianxi-blog/service/blog/model"

type (
	PostWrapper struct {
		Post      *model.Posts     `json:"post"`
		NextTitle string           `json:"next_title"`
		PreTitle  string           `json:"pre_title"`
		Comments  []model.Comments `json:"comments"`
	}
)
