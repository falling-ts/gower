package services

type Passwd interface {
	Hash(passwd string) (string, error)
	Check(passwd string, hash string) error
}

type PasswdService interface {
	Service

	Hash(passwd string) (string, error)
	Check(passwd string, hash string) error
}
