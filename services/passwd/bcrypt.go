package passwd

import "golang.org/x/crypto/bcrypt"

type _bcrypt struct{}

// Hash 哈希加密
func (b *_bcrypt) Hash(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// Check 校验密码
func (b *_bcrypt) Check(passwd string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwd))
}
