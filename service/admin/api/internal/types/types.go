// Code generated by goctl. DO NOT EDIT.
package types

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}