package controllers

// Response template 成功响应体
type Response struct {
	HttpStatus uint
	Name       string
	Data       map[string]any
}

// JsonResponse json 成功响应体
type JsonResponse struct {
	Code uint `json:"code"`
	Msg  uint `json:"msg"`
	Data any  `json:"data"`
}

type Controllers struct{}

type data map[string]any
