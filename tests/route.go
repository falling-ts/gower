package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestControllers(t *testing.T) {
	assert := getAssert(t)
	route.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"Test": "123",
		})
	})
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)
}
