package util

import (
	"crypto/rand"
	"encoding/base64"
)

// SecretKey 加密安全的伪随机数生成器
func (*Service) SecretKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(randomBytes), nil
}
