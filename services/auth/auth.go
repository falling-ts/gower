package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"gower/services"
	"gower/utils/slice"

	"github.com/golang-jwt/jwt/v5"
)

// Service Auth 服务
type Service struct {
	*jwt.Token
}

// New 新建 Auth 服务
func New() *Service {
	return new(Service)
}

var (
	config services.Config
	util   services.UtilService
	cache  services.CacheService
)

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	util = args[1].(services.UtilService)
	cache = args[2].(services.CacheService)

	return s
}

// Sign 签名 Token
func (s *Service) Sign(args ...any) (string, error) {
	var method *jwt.SigningMethodHMAC
	method = map[string]*jwt.SigningMethodHMAC{
		"HS256": jwt.SigningMethodHS256,
		"HS384": jwt.SigningMethodHS384,
		"HS512": jwt.SigningMethodHS512,
	}[config.Get("jwt.method", "HS256").(string)]
	if method == nil {
		method = jwt.SigningMethodHS256
	}

	claims := new(Claims).Set(args...)
	if claims.UpdateDur == nil {
		updateDur := config.Get("jwt.upd", 5*time.Minute).(time.Duration)
		claims.Set(jwt.NewNumericDate(time.Now().Add(updateDur)))
	}
	if claims.ExpiresAt == nil {
		exp := config.Get("jwt.exp", 10*time.Minute).(time.Duration)
		claims.Set(jwt.NewNumericDate(time.Now().Add(exp)))
	}

	token := jwt.NewWithClaims(method, *claims.Set(util.Nanoid()))

	key, err := base64.StdEncoding.DecodeString(config.Get("jwt.key").(string))
	if err != nil {
		return "", err
	}
	return token.SignedString(key)
}

// Check 校验 Token
func (s *Service) Check(token string, args ...string) (string, string, error) {
	claims, err := s.parseClaims(token)
	if err != nil {
		return "", "", err
	}

	for _, arg := range args {
		if !s.checkAudience(arg, claims.Audience) {
			s.execBlack(claims.ID, claims.ExpiresAt.Sub(time.Now()))
			return "", "", errors.New("token 身份存疑, 已拉黑")
		}
	}

	if claims.UpdateDur.Before(time.Now()) {
		var newToken string
		newToken, err = s.Sign(claims.Issuer, claims.Subject, claims.Audience)
		if err != nil {
			return "", "", err
		}

		return claims.Subject, newToken, nil
	}

	return claims.Subject, "", nil
}

// Black 拉黑 Token
func (s *Service) Black(token string) error {
	claims, err := s.parseClaims(token)
	if err != nil {
		return err
	}
	s.execBlack(claims.ID, claims.ExpiresAt.Sub(time.Now()))
	return nil
}

// IsToken 判断是否是 Token
func (s *Service) IsToken(token string) bool {
	tokens := strings.Split(token, ".")
	if len(tokens) < 3 {
		return false
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(tokens[0])
	if err != nil {
		return false
	}
	header := make(map[string]any)
	err = json.Unmarshal(headerBytes, &header)
	if err != nil {
		return false
	}

	if typ, ok := header["typ"]; !ok || typ != "JWT" {
		return false
	}

	return true
}

func (s *Service) parseClaims(t string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(t, new(Claims), func(token *jwt.Token) (interface{}, error) {
		return base64.StdEncoding.DecodeString(config.Get("jwt.key").(string))
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token 无效")
	}

	id := claims.ID
	if id == "" {
		return nil, errors.New("token id 不能为空")
	}

	_, ok = cache.Get(util.BlackTokenKey(id))
	if ok {
		return nil, errors.New("token 已拉黑")
	}

	return claims, nil
}

func (s *Service) checkAudience(str string, aud jwt.ClaimStrings) bool {
	return slice.Strings(aud).Has(str)
}

func (s *Service) execBlack(id string, d time.Duration) {
	cache.Set(util.BlackTokenKey(id), struct{}{}, d)
}
