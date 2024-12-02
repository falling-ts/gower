package services

import "reflect"

type UtilService interface {
	Service

	Nanoid(args ...int) string
	Direct(v reflect.Value) reflect.Value

	SetEnv(env, key, value string) error
	SecretKey(length int) (string, error)

	ExcKey() string
	BlackTokenKey(nanoid string) string

	Ptr(v any) any

	CreateDir(dir string) string
	IsExist(file string) bool

	SHA256(str string) string
}
