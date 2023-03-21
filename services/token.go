package services

import "time"

type TokenService interface {
	Service

	Sign(args ...any) (string, error)
	Check(token string, args ...string) (string, string, error)
	Block(id string, d time.Duration)
	IsToken(token string) bool
}
