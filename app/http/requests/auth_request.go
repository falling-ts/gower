package requests

import "gower/app"

type AuthRequest struct {
	app.Request
	Password string `form:"password" json:"password" binding:"required" zh:"密码"`
}

type RegisterRequest struct {
	AuthRequest
	Username *string `form:"username" json:"username" binding:"required,alphanum" zh:"用户名"`
	Email    *string `form:"email" json:"email" binding:"required,email" zh:"邮箱"`
}

type LoginRequest struct {
	AuthRequest
	Username *string `form:"username" json:"username" binding:"required" zh:"账户"`
}
