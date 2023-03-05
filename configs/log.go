package configs

type Log struct {
	SkipPaths []string `env:"SKIP_PATHS" envSeparator:","`
	Channel   string   `env:"LOG_CHANNEL" envDefault:"stack"`
	Dir       string   `env:"LOG_DIR" envDefault:"storage/logs"`
}
