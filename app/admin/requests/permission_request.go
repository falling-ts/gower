package requests

import "gitee.com/falling-ts/gower/app"

type PermissionRequest struct {
	app.Request

	Name *string `form:"name" json:"name" binding:"required"`
}
