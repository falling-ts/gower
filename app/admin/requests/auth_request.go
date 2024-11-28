package requests

import "gitee.com/falling-ts/gower/app"

type AuthRequest struct {
	app.Request

	Username *string `form:"username" json:"username" binding:"required" zh:"账户"`
	Password string  `form:"password" json:"password" binding:"required" zh:"密码"`
}
