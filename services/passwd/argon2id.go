package passwd

import (
	"github.com/alexedwards/argon2id"
	"net/http"
)

type _argon2id struct{}

// Hash 哈希加密
func (a *_argon2id) Hash(passwd string) (string, error) {
	return argon2id.CreateHash(passwd, argon2id.DefaultParams)
}

// Check 校验密码
func (a *_argon2id) Check(passwd string, hash string) error {
	match, err := argon2id.ComparePasswordAndHash(passwd, hash)
	if err != nil {
		return err
	}

	if !match {
		return exception.New(http.StatusUnauthorized, "密码错误.")
	}

	return nil
}
