package response

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/falling-ts/gower/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 接受的响应数据类型
var accepts = []string{
	gin.MIMEJSON,
	gin.MIMEHTML,
	gin.MIMEXML,
	gin.MIMEYAML,
	gin.MIMETOML,
	gin.MIMEPlain,
}

// Service 响应结构体
type Service struct {
	services.Response
	HttpStatus int
	config     gin.Negotiate
}

var (
	auth   services.AuthService
	cookie services.CookieService
	util   services.UtilService
	config services.Config
	db     services.DBService
	cache  services.CacheService
	excp   services.Exception
)

// Mount 挂载响应体
func Mount(r services.Response) services.Response {
	return r.Set(new(Service)).Set(r)
}

// New 新建响应服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	auth = args[0].(services.AuthService)
	cookie = args[1].(services.CookieService)
	util = args[2].(services.UtilService)
	config = args[3].(services.Config)
	db = args[4].(services.DBService)
	cache = args[5].(services.CacheService)
	excp = args[6].(services.Exception)
	return s.Response
}

// Build 构建每个请求的异常
func (s *Service) Build(code int, args ...any) services.Response {
	s.config.Offered = config.Get("res.mimes", accepts).([]string)

	s.decideType("success")
	argsNum := len(args)
	for i := 0; i < argsNum; i++ {
		s.decideType(args[i])
	}

	s.HttpStatus = code

	return s.Response
}

// Handle 处理响应
func (s *Service) Handle(c *gin.Context) bool {
	s.bodyLogger(c)
	s.handleToken(c)
	s.csrfTokenAndCommonData(c)
	s.adminData(c)
	c.Negotiate(s.HttpStatus, s.config)
	return true
}

// IsToken 判断是否是 Token
func (s *Service) IsToken(token string) bool {
	return auth.IsToken(token)
}

func (s *Service) decideType(arg any) {
	switch arg.(type) {
	case int:
		code := arg.(int)
		if code >= http.StatusOK && code < http.StatusMultipleChoices {
			s.HttpStatus = code
		} else {
			s.Response.Set(arg)
		}
	case string:
		s.Response.Set(arg)
		s.config.HTMLName = arg.(string)
	default:
		s.Response.Set(arg)
		s.config.HTMLData = arg
	}

	s.config.Data = s.Response
}

func (s *Service) bodyLogger(c *gin.Context) {
	if c.NegotiateFormat(s.config.Offered...) != gin.MIMEHTML {
		c.Set("body-logger", s.Response)
	} else {
		c.Set("body-logger", "html body")
	}
}

func (s *Service) handleToken(c *gin.Context) {
	tokenKey := c.GetString("token-key")
	if tokenKey == "" {
		tokenKey = "token"
	}

	token := c.GetString(tokenKey)
	if token == "" {
		tokenAny, _ := s.Response.Get("token")
		token = tokenAny.(string)
	}

	if token != "" {
		cookie.Set(c, tokenKey, token, false)
		s.Set(token)
	}
}

func (s *Service) csrfTokenAndCommonData(c *gin.Context) {
	mime := c.NegotiateFormat(s.config.Offered...)
	if mime == binding.MIMEHTML {
		data := reflect.Indirect(reflect.ValueOf(s.config.HTMLData))
		if data.Kind() == reflect.Map {
			csrfToken := util.Nanoid()
			data.SetMapIndex(reflect.ValueOf("csrf_token"), reflect.ValueOf(csrfToken))
			cookie.Set(c, "csrf_token", csrfToken)

			titleKey := "app_title"
			titleVal := data.MapIndex(reflect.ValueOf(titleKey))
			if !titleVal.IsValid() {
				title := config.Get("app.name", "Gower").(string)
				data.SetMapIndex(reflect.ValueOf(titleKey), reflect.ValueOf(title))
			}

			themeKey := "app_theme"
			themeVal := data.MapIndex(reflect.ValueOf(themeKey))
			if !themeVal.IsValid() {
				theme := config.Get("view.theme", "lofi").(string)
				data.SetMapIndex(reflect.ValueOf(themeKey), reflect.ValueOf(theme))
			}

			excpKey, err := cookie.Get(c, "exception")
			if err == nil {
				exceptionKey := "app_exceptions"
				exceptionVal := data.MapIndex(reflect.ValueOf(exceptionKey))
				if !exceptionVal.IsValid() {
					if exception, ok := cache.Flash(excpKey); ok {
						data.SetMapIndex(reflect.ValueOf(exceptionKey), reflect.ValueOf(exception))
					}
				}
			}
		}
	}
}

func (s *Service) adminData(c *gin.Context) {
	mime := c.NegotiateFormat(s.config.Offered...)
	if mime == binding.MIMEHTML && strings.HasPrefix(c.FullPath(), "/admin") {
		data := reflect.Indirect(reflect.ValueOf(s.config.HTMLData))
		if data.Kind() == reflect.Map {
			menusKey := "admin_menus"
			menusVal := data.MapIndex(reflect.ValueOf(menusKey))
			if !menusVal.IsValid() {
				if menus, ok := c.Get(menusKey); ok {
					data.SetMapIndex(reflect.ValueOf(menusKey), reflect.ValueOf(menus))
				}
			}

		}
	}
}
