package configs

import "time"

type Cache struct {
	Expire   time.Duration `env:"CACHE_EXPIRE" envDefault:"5m"`
	Clean    time.Duration `env:"CACHE_CLEAN" envDefault:"10m"`
	Interval time.Duration `env:"CACHE_INTERVAL" envDefault:"10m"`
	Dir      string        `env:"CACHE_DIR" envDefault:"storage/caches"`
	FILE     string        `env:"CACHE_FILE" envDefault:"go.cache"`
}
