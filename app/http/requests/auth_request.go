package requests

type AuthRequest struct {
	*request
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterRequest struct {
	AuthRequest
}

type LoginRequest struct {
	AuthRequest
}
