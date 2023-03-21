package services

type UtilService interface {
	Service

	Nanoid(args ...int) string

	SetEnv(key, value string) error
	SecretKey(length int) (string, error)

	ExcpKey() string
	BlackTokenKey(nanoid string) string
}
