package configs

import "time"

type Jwt struct {
	Key    string        `env:"JWT_KEY,required"`
	Upd    time.Duration `env:"JWT_UPD" envDefault:"5m"`
	Exp    time.Duration `env:"JWT_EXP" envDefault:"10m"`
	Method string        `env:"JWT_METHOD" envDefault:"HS256"`
}
