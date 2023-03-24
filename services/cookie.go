package services

import "github.com/gin-gonic/gin"

type CookieService interface {
	Service

	Set(c *gin.Context, key, val string, args ...any)
	Get(c *gin.Context, key string) (string, error)
}
