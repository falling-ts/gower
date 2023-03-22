package services

type AuthService interface {
	Service

	Sign(args ...any) (string, error)
	Check(token string, args ...string) (string, string, error)
	Black(token string) error
	IsToken(token string) bool
}
