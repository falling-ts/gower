package services

import (
	"github.com/gin-gonic/gin"
)

type Storage interface {
	Image(c *gin.Context) (string, string, error)
	File(c *gin.Context) (string, string, error)
}

type UploadService interface {
	Service

	Image(c *gin.Context) (string, string, error)
	File(c *gin.Context) (string, string, error)
	Store(storage string) Storage
}
