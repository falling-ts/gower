package route

import (
	"bufio"
	"encoding/json"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	ctx *gin.Context
}

var _ gin.ResponseWriter = (*responseWriter)(nil)

func setWriter(c *gin.Context) {
	writer := c.Writer
	c.Writer = &responseWriter{
		ResponseWriter: writer,
		ctx:            c,
	}
}

func (w *responseWriter) Unwrap() http.ResponseWriter {
	return w.ResponseWriter
}

func (w *responseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) WriteHeaderNow() {
	w.ResponseWriter.WriteHeaderNow()
}

func (w *responseWriter) Write(data []byte) (n int, err error) {
	if json.Valid(data) {
		w.ctx.Set("body-logger", string(data))
	}
	n, err = w.ResponseWriter.Write(data)
	return
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
	n, err = w.ResponseWriter.WriteString(s)
	return
}

func (w *responseWriter) Status() int {
	return w.ResponseWriter.Status()
}

func (w *responseWriter) Size() int {
	return w.ResponseWriter.Size()
}

func (w *responseWriter) Written() bool {
	return w.ResponseWriter.Written()
}

// Hijack implements the http.Hijacker interface.
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.Hijack()
}

// CloseNotify implements the http.CloseNotifier interface.
func (w *responseWriter) CloseNotify() <-chan bool {
	return w.ResponseWriter.CloseNotify()
}

// Flush implements the http.Flusher interface.
func (w *responseWriter) Flush() {
	w.ResponseWriter.Flush()
}

func (w *responseWriter) Pusher() (pusher http.Pusher) {
	return w.ResponseWriter.Pusher()
}
