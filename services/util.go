package services

import "reflect"

type UtilService interface {
	Service

	Nanoid(args ...int) string
	Direct(v reflect.Value) reflect.Value

	SetEnv(key, value string) error
	SecretKey(length int) (string, error)

	ExcpKey() string
	BlackTokenKey(nanoid string) string
}
