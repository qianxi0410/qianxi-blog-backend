package util

type Code int

// 返回的消息体
type Response struct {
	Code Code        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 状态码(自定义)
const OK Code = 666
const ERROR Code = 777
