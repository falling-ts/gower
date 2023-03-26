package configs

type Passwd struct {
	Mode string `env:"PASSWD_MODE" envDefault:"argon2id"`
}
