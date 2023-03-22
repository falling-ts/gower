package configs

import "time"

type DB struct {
	Driver            string        `env:"DB_DRIVER" envDefault:"mysql"`
	User              string        `env:"DB_USER" envDefault:"root"`
	Passwd            string        `env:"DB_PASSWD" envDefault:"root"`
	Net               string        `env:"DB_NET" envDefault:"tcp"`
	Host              string        `env:"DB_HOST" envDefault:"localhost"`
	Port              int           `env:"DB_PORT" envDefault:"3306"`
	Name              string        `env:"DB_NAME" envDefault:"gower"`
	MaxOpen           int           `env:"DB_MAX_OPEN" envDefault:"100"`
	MaxIdleCount      int           `env:"DB_MAX_IDLE_COUNT" envDefault:"25"`
	MaxLifeTime       time.Duration `env:"DB_MAX_LIFE_TIME" envDefault:"30m"`
	MaxIdleTime       time.Duration `env:"DB_MAX_IDLE_TIME" envDefault:"10m"`
	DisableForeignKey bool          `env:"DB_DISABLE_FOREIGN_KEY" envDefault:"true"`
	Timezone          string        `env:"DB_TIMEZONE" envDefault:"sys"`
	Mysql             struct {
		AllowNativePasswords bool `env:"DB_MYSQL_NATIVE_PASSWORDS" envDefault:"true"`
	}
}
