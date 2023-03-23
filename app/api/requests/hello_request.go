package requests

import "gower/app"

type HelloRequest struct {
	app.Request
	Test string `form:"name" binding:"required" zh:"测试"`
}

type IndexRequest struct {
	HelloRequest
	Name *string `form:"name" binding:"required" zh:"名字"`
}
