package cookie

import (
	"github.com/falling-ts/gower/services"
	"github.com/gin-gonic/gin"
)

type Service struct {
	maxAge   int
	path     string
	domain   string
	httpOnly bool
	secure   bool
}

var (
	config   services.Config
	symCrypt services.SymCryptService
)

// New 新建服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	symCrypt = args[1].(services.SymCryptService)

	s.maxAge = 100000000
	s.path = "/"
	s.domain = config.Get("app.domain", "localhost").(string)
	s.httpOnly = true
	s.secure = false

	return s
}

// Set 加密设置 Cookie
func (s *Service) Set(c *gin.Context, key, val string, args ...any) {
	newS := s.clone()
	for i := 0; i < len(args); i++ {
		newS.decideType(s, args[i])
	}

	cryVal, _ := symCrypt.Encrypt(val)

	c.SetCookie(key, cryVal, newS.maxAge, newS.path, newS.domain, newS.secure, newS.httpOnly)
}

// Get 解密获取 Cookie
func (s *Service) Get(c *gin.Context, key string) (string, error) {
	val, err := c.Cookie(key)
	if err != nil {
		return val, err
	}

	return symCrypt.Decrypt(val)
}

func (s *Service) clone() *Service {
	temp := *s
	return &temp
}

func (s *Service) decideType(proto *Service, arg any) {
	switch arg.(type) {
	case string:
		str := arg.(string)
		if str != proto.path {
			s.path = str
			break
		}
		if str != proto.domain {
			s.domain = str
		}
	case int:
		s.maxAge = arg.(int)
	case bool:
		boolean := arg.(bool)
		if boolean != proto.httpOnly {
			s.httpOnly = boolean
			break
		}
		if boolean != proto.secure {
			s.secure = boolean
		}
	}
}
