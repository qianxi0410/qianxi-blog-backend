package util

// 返回的消息体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 状态码
const OK = 200
const NOT_FOUND = 404
