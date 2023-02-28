package controllers

import "gower/app/services/route"

type PongController struct {
	Controllers
}

var Pong = new(PongController)

func (p *PongController) Ping(c route.Context) {
	c.JSON(200, map[string]any{
		"message": "pong",
	})
}
