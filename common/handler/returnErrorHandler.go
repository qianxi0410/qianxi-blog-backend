package handler

import "net/http"

// 返回的错误处理
// 策略是无论服务端处理如何, 都返回200
// 将错误消息包装在json对象中

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnErrorHandler(err error) (int, interface{}) {
	return http.StatusOK, response{
		Code: 777,
		Msg:  err.Error(),
		Data: nil,
	}
}
