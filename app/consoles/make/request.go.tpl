package requests

import "{{.Module}}/app"

type {{.UpCamel}}Request struct {
	app.Request

	Name *string `form:"name" json:"name" binding:"required"`
}
