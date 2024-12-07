package requests

import "gitee.com/falling-ts/gower/app"

type AdminRequest struct {
	app.Request
	Email    string `form:"email" json:"email"`
	Nickname string `form:"nickname" json:"nickname"`
	Avatar   string `form:"avatar" json:"avatar"`
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

type AdminStoreRequest struct {
	AdminRequest
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
