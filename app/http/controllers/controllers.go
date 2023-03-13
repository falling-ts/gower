package controllers

type Controllers struct{}

type data map[string]any

// Response template 成功响应体
type Response struct {
	HttpStatus int
	Name       string
	Data       map[string]any
}

// JsonResponse json 成功响应体
type JsonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
