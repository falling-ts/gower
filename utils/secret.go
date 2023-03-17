package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// SecretKey 生成安全秘钥
func SecretKey(length int) (string, error) {
	// 确保 Base64 编码后的字符串长度至少为 length
	randomBytes := make([]byte, (length*6+7)/8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(randomBytes)
	return encoded[:length], nil
}
