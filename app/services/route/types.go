package route

import (
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/render"
)

// Service route service.
type Service interface {
	Run()
}

type ErrorType uint64

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(Context)

// HandlersChain defines a HandlerFunc slice.
type HandlersChain []HandlerFunc

// Context http request and response.
type Context interface {
	Copy() *gin.Context
	HandlerName() string
	HandlerNames() []string
	Handler() gin.HandlerFunc
	FullPath() string

	Next()
	IsAborted() bool
	Abort()
	AbortWithStatus(code int)
	AbortWithStatusJSON(code int, jsonObj any)
	AbortWithError(code int, err error) *gin.Error
	Error(err error) *gin.Error

	Set(key string, value any)
	Get(key string) (value any, exists bool)
	MustGet(key string) any
	GetString(key string) (s string)
	GetBool(key string) (b bool)
	GetInt(key string) (i int)
	GetInt64(key string) (i64 int64)
	GetUint(key string) (ui uint)
	GetUint64(key string) (ui64 uint64)
	GetFloat64(key string) (f64 float64)
	GetTime(key string) (t time.Time)
	GetDuration(key string) (d time.Duration)
	GetStringSlice(key string) (ss []string)
	GetStringMap(key string) (sm map[string]any)
	GetStringMapString(key string) (sms map[string]string)
	GetStringMapStringSlice(key string) (smss map[string][]string)

	Param(key string) string
	AddParam(key, value string)
	Query(key string) (value string)
	DefaultQuery(key, defaultValue string) string
	GetQuery(key string) (string, bool)
	QueryArray(key string) (values []string)
	GetQueryArray(key string) (values []string, ok bool)
	QueryMap(key string) (dicts map[string]string)
	GetQueryMap(key string) (map[string]string, bool)
	PostForm(key string) (value string)
	DefaultPostForm(key, defaultValue string) string
	GetPostForm(key string) (string, bool)
	PostFormArray(key string) (values []string)
	GetPostFormArray(key string) (values []string, ok bool)
	PostFormMap(key string) (dicts map[string]string)
	GetPostFormMap(key string) (map[string]string, bool)
	FormFile(name string) (*multipart.FileHeader, error)
	MultipartForm() (*multipart.Form, error)
	SaveUploadedFile(file *multipart.FileHeader, dst string) error

	Bind(obj any) error
	BindJSON(obj any) error
	BindXML(obj any) error
	BindQuery(obj any) error
	BindYAML(obj any) error
	BindTOML(obj interface{}) error
	BindHeader(obj any) error
	BindUri(obj any) error
	MustBindWith(obj any, b binding.Binding) error
	ShouldBind(obj any) error
	ShouldBindJSON(obj any) error
	ShouldBindXML(obj any) error
	ShouldBindQuery(obj any) error
	ShouldBindYAML(obj any) error
	ShouldBindTOML(obj interface{}) error
	ShouldBindHeader(obj any) error
	ShouldBindUri(obj any) error
	ShouldBindWith(obj any, b binding.Binding) error
	ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error)

	ClientIP() string
	RemoteIP() string
	ContentType() string
	IsWebsocket() bool

	Status(code int)
	Header(key, value string)
	GetHeader(key string) string
	GetRawData() ([]byte, error)
	SetSameSite(samesite http.SameSite)
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	Cookie(name string) (string, error)

	Render(code int, r render.Render)
	HTML(code int, name string, obj any)
	IndentedJSON(code int, obj any)
	SecureJSON(code int, obj any)
	JSONP(code int, obj any)
	JSON(code int, obj any)
	AsciiJSON(code int, obj any)
	PureJSON(code int, obj any)
	XML(code int, obj any)
	YAML(code int, obj any)
	TOML(code int, obj interface{})
	ProtoBuf(code int, obj any)
	String(code int, format string, values ...any)
	Redirect(code int, location string)

	Data(code int, contentType string, data []byte)
	DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string)

	File(filepath string)
	FileFromFS(filepath string, fs http.FileSystem)
	FileAttachment(filepath, filename string)

	SSEvent(name string, message any)
	Stream(step func(w io.Writer) bool) bool

	Negotiate(code int, config gin.Negotiate)
	NegotiateFormat(offered ...string) string

	SetAccepted(formats ...string)

	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}
