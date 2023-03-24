package services

type SymCryptService interface {
	Service

	Encrypt(plain string) (string, error)
	Decrypt(cipher string) (string, error)
}
