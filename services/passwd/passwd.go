package passwd

import (
	"errors"
	"gitee.com/falling-ts/gower/services"
)

type Service struct {
	services.Passwd
}

var (
	config    services.Config
	exception services.Exception
)

func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	config = args[0].(services.Config)
	errors.As(args[1].(services.Exception), &exception)

	var ok bool
	s.Passwd, ok = map[string]services.Passwd{
		"bcrypt":   new(_bcrypt),
		"argon2id": new(_argon2id),
		"scrypt":   new(_scrypt),
	}[config.Get("passwd.mode", "argon2id").(string)]
	if !ok {
		s.Passwd = new(_bcrypt)
	}

	return s
}
