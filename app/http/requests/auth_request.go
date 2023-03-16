package requests

type AuthRequest struct {
	request
	Username string `form:"username" json:"username" binding:"required,alphanum" zh:"用户名"`
	Password string `form:"password" json:"password" binding:"required" zh:"密码"`
}

type RegisterRequest struct {
	AuthRequest
}

type LoginRequest struct {
	AuthRequest
}
