package symcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/falling-ts/gower/services"
	"io"
)

type Service struct {
	key []byte
}

var config services.Config

// New 新建服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(args ...services.Service) services.Service {
	var err error
	config = args[0].(services.Config)
	base64Key := config.Get("app.key").(string)
	s.key, err = base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		panic(err)
	}

	return s
}

// Encrypt 加密
func (s *Service) Encrypt(plaintext string) (string, error) {
	// 创建一个新的 AES 加密块，使用给定的密钥
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return plaintext, err
	}

	// 对明文进行补全
	blockSize := block.BlockSize()
	plainBytes := pkcs7Padding([]byte(plaintext), blockSize)

	// 创建一个 AES CBC 加密器，使用随机生成的 IV
	ciphertext := make([]byte, blockSize+len(plainBytes))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return plaintext, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)

	// 加密
	mode.CryptBlocks(ciphertext[blockSize:], plainBytes)

	// 返回加密结果，使用 Base64 编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密
func (s *Service) Decrypt(ciphertext string) (string, error) {
	// 解码 Base64
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return ciphertext, err
	}

	// 创建一个新的 AES 加密块，使用给定的密钥
	block, err := aes.NewCipher(s.key)
	if err != nil {
		return ciphertext, err
	}

	// 创建一个 AES CBC 解密器，使用密文中的 IV
	blockSize := block.BlockSize()
	if len(ciphertextBytes) < blockSize {
		return ciphertext, errors.New("密文错误")
	}
	iv := ciphertextBytes[:blockSize]
	ciphertextBytes = ciphertextBytes[blockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密
	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	// 去除填充的字节
	return string(unPkcs7Padding(ciphertextBytes)), nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unPkcs7Padding(data []byte) []byte {
	length := len(data)
	paddingLen := int(data[length-1])
	return data[:(length - paddingLen)]
}
