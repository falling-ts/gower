package util

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/falling-ts/gower/services"
	"github.com/jaevor/go-nanoid"
	"os"
	"reflect"
	"strings"
)

type Service struct{}

// New 新建 Util 服务
func New() *Service {
	return new(Service)
}

// Init 初始化
func (s *Service) Init(...services.Service) services.Service {
	return s
}

// Nanoid 获取简单唯一 ID
func (s *Service) Nanoid(args ...int) string {
	arg := 21
	if len(args) > 0 {
		arg = args[0]
	}

	genKey, err := nanoid.Standard(arg)
	if err != nil {
		panic(err)
	}
	return genKey()
}

// Direct 获取反射指针类型
func (s *Service) Direct(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Pointer {
		return v.Addr()
	}

	return v
}

// SecretKey 加密安全的伪随机数生成器
func (*Service) SecretKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(randomBytes), nil
}

// SetEnv 设置 env
func (*Service) SetEnv(env, key, value string) error {
	file, err := os.Open(env)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			line = fmt.Sprintf("%s=%s", key, value)
		}
		lines = append(lines, line)
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	outputFile, err := os.Create(env)
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		_ = outputFile.Close()
	}(outputFile)

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// Ptr 获取类型的指针类型
func (s *Service) Ptr(v any) any {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Bool:
		val := v.(bool)
		return &val
	case reflect.Int:
		val := v.(int)
		return &val
	case reflect.Int8:
		val := v.(int8)
		return &val
	case reflect.Int16:
		val := v.(int16)
		return &val
	case reflect.Int32:
		val := v.(int32)
		return &val
	case reflect.Int64:
		val := v.(int64)
		return &val
	case reflect.Uint:
		val := v.(uint)
		return &val
	case reflect.Uint8:
		val := v.(uint8)
		return &val
	case reflect.Uint16:
		val := v.(uint16)
		return &val
	case reflect.Uint32:
		val := v.(uint32)
		return &val
	case reflect.Uint64:
		val := v.(uint64)
		return &val
	case reflect.Uintptr:
		val := v.(uintptr)
		return &val
	case reflect.Float32:
		val := v.(float32)
		return &val
	case reflect.Float64:
		val := v.(float64)
		return &val
	case reflect.Complex64:
		val := v.(complex64)
		return &val
	case reflect.Complex128:
		val := v.(complex128)
		return &val
	case reflect.String:
		val := v.(string)
		return &val
	}

	return v
}

// CreateDir 如果目录不存在, 则创建
func (s *Service) CreateDir(dir string) string {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}

	return dir
}

// IsExist 判断文件是否存在
func (s *Service) IsExist(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// SHA256 哈希计算
func (s *Service) SHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}
