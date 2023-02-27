package controllers

import "github.com/gin-gonic/gin"

type PongController struct {
	Controllers
}

var Pong = new(PongController)

func (p *PongController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
