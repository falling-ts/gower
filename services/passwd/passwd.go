package passwd

import "gower/services"

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
func (s *Service) Init(args ...any) {
	config = args[0].(services.Config)
	exception = args[1].(services.Exception)

	switch config.Get("passwd.mode", "bcrypt").(string) {
	case "argon2id":
		s.Passwd = new(_argon2id)
	case "scrypt":
		s.Passwd = new(_scrypt)
	default:
		s.Passwd = new(_bcrypt)
	}
}
