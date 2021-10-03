// Code generated by goctl. DO NOT EDIT.
package types

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CountWithTagReq struct {
	Tag string `path:"tag"`
}

type PageReq struct {
	Page int64 `path:"page"`
	Size int64 `path:"size"`
}

type PageWithTagReq struct {
	Page int64  `path:"page"`
	Size int64  `path:"size"`
	Tag  string `path:"tag"`
}

type PostReq struct {
	Id int64 `path:"id"`
}

type DeleteReq struct {
	Id    int64  `path:"id"`
	Login string `path:"login"`
}

type SaveReq struct {
	Content string `json:"content"`
	Login   string `json:"login"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	PostId  int64  `json:"post_id"`
}

type Oauth2Req struct {
	Code string `path:"code"`
}
