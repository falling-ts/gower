package passwd

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"net/http"
	"strings"
)

type _scrypt struct{}

// Hash 哈希加密
func (i *_scrypt) Hash(passwd string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash, err := scrypt.Key([]byte(passwd), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	return strings.Join([]string{encodedHash, encodedSalt}, "."), err
}

// Check 校验密码
func (i *_scrypt) Check(passwd string, hash string) error {
	hashes := strings.Split(hash, ".")
	pwdHash, _ := base64.StdEncoding.DecodeString(hashes[0])
	slat, _ := base64.StdEncoding.DecodeString(hashes[1])

	inputHash, err := scrypt.Key([]byte(passwd), slat, 32768, 8, 1, 32)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare(pwdHash, inputHash) == 1 {
		return nil
	} else {
		return exception.New(http.StatusUnauthorized, "密码错误.")
	}
}
