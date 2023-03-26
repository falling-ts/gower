package requests

import "gower/app"

type TestRequest struct {
	app.Request

	Test string `form:"test" json:"test" binding:"required"`
}
