package requests

import "gitee.com/falling-ts/gower/app"

type AdminRequest struct {
	app.Request
	Name *string `form:"name" json:"name" binding:"required"`
}

type AdminIndexRequest struct {
	app.IndexRequest
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Nickname string `form:"nickname" json:"nickname"`
}
