package requests

import "github.com/falling-ts/gower/app"

type {{.UpCamel}}Request struct {
	app.Request

	Name *string `form:"name" json:"name" binding:"required"`
}
