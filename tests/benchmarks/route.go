package benchmarks

import (
	"github.com/falling-ts/gower/app/http/requests"
	"github.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type data map[string]any

func init() {
	route.GET("/bench01", func(c *gin.Context) {
		c.JSON(http.StatusOK, data{
			"benchmark": "benchmark",
		})
	})
	route.GET("/bench02", func(c *gin.Context) error {
		c.JSON(http.StatusOK, data{
			"benchmark": "benchmark",
		})
		return nil
	})
	route.GET("/bench03", func(c *gin.Context) services.Response {
		return res.Ok("请求成功", data{
			"benchmark": "benchmark",
		})
	})
	route.GET("/bench04", func(*gin.Context) (services.Response, error) {
		return res.Ok("请求成功", data{
			"benchmark": "benchmark",
		}), nil
	})
	route.GET("/bench05", func() error {
		return nil
	})
	route.GET("/bench06", func() services.Response {
		return res.Ok("请求成功", data{
			"benchmark": "benchmark",
		})
	})
	route.GET("/bench07", func() (services.Response, error) {
		return res.Ok("请求成功", data{
			"benchmark": "benchmark",
		}), nil
	})
	route.GET("/bench08", "请求成功")
	route.GET("/bench09", []any{"请求成功", data{"benchmark": "benchmark"}})
	route.GET("/bench10", func(req *requests.TestRequest) (string, any) {
		return "请求成功", data{"benchmark": req.Test}
	})
}

func BenchmarkRoute01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench01", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench02", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench03", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute04(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench04", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute05(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench05", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute06(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench06", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute07(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench07", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute08(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench08", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench09", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}

func BenchmarkRoute10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/bench10?test=benchmark", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
	}
}
