package requests

import "github.com/falling-ts/gower/app"

type TestRequest struct {
	app.Request

	Test *string `form:"test" json:"test" binding:"required"`
}
