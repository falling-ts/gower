package configs

import "time"

type Cors struct {
	AllowOrigins  []string      `env:"CORS_ALLOW_ORIGINS" envSeparator:"," envDefault:"*"`
	AllowMethods  []string      `env:"CORS_ALLOW_METHODS" envSeparator:"," envDefault:"*"`
	AllowHeaders  []string      `env:"CORS_ALLOW_HEADERS" envSeparator:"," envDefault:"*"`
	ExposeHeaders []string      `env:"CORS_EXPOSE_HEADERS" envSeparator:"," envDefault:"*"`
	MaxAge        time.Duration `env:"CORS_MAX_AGE" envDefault:"12h"`
}
