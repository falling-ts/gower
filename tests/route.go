package tests

import (
	"fmt"
	"gower/app/http/requests"
	"gower/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type data map[string]any

func TestRoute(t *testing.T) {
	fmt.Println("----------------TestResponse 开始----------------")

	assert := getAssert(t)
	var (
		req *http.Request
		w   *httptest.ResponseRecorder
	)

	route.GET("/test01", func(c *gin.Context) {
		c.JSON(http.StatusOK, data{
			"id": "1",
		})
	})
	req = httptest.NewRequest(http.MethodGet, "/test01", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test02", func(c *gin.Context) error {
		c.JSON(http.StatusOK, data{
			"id": "1",
		})
		return nil
	})
	req = httptest.NewRequest(http.MethodGet, "/test02", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test03", func(c *gin.Context) services.Response {
		return res.Ok("请求成功")
	})
	req = httptest.NewRequest(http.MethodGet, "/test03", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test04", func(*gin.Context) (services.Response, error) {
		return res.Ok("请求成功"), nil
	})
	req = httptest.NewRequest(http.MethodGet, "/test04", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test05", func() error {
		return nil
	})
	req = httptest.NewRequest(http.MethodGet, "/test05", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test06", func() services.Response {
		return res.Ok("请求成功")
	})
	req = httptest.NewRequest(http.MethodGet, "/test06", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test07", func() (services.Response, error) {
		return res.Ok("请求成功"), nil
	})
	req = httptest.NewRequest(http.MethodGet, "/test07", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test08", "请求成功")
	req = httptest.NewRequest(http.MethodGet, "/test08", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test09", []any{"请求成功", data{"id": 1}})
	req = httptest.NewRequest(http.MethodGet, "/test09", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test10", []any{"请求成功", data{"id": 1}})
	req = httptest.NewRequest(http.MethodGet, "/test10", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	route.GET("/test11", func(req *requests.TestRequest) (string, any) {
		return "请求成功", data{"test": req.Test}
	})
	req = httptest.NewRequest(http.MethodGet, "/test11?test=hello", nil)
	w = httptest.NewRecorder()
	route.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)

	fmt.Println("----------------TestResponse 结束----------------")
}
