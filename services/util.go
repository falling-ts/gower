package services

type UtilService interface {
	Service

	SetEnv(key, value string) error
	SecretKey(length int) (string, error)
}
