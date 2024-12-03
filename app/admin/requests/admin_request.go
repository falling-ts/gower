package requests

import "gitee.com/falling-ts/gower/app"

type AdminRequest struct {
	app.Request
}

type AdminIndexRequest struct {
	app.IndexRequest
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Nickname string `form:"nickname" json:"nickname"`
}

type AdminCreateRequest struct {
	app.ModalRequest
}
