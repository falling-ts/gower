package configs

type Log struct {
	Dir            string   `env:"LOG_DIR" envDefault:"storage/logs"`
	Channel        string   `env:"LOG_CHANNEL" envDefault:"stack"`
	SkipPaths      []string `env:"LOG_SKIP_PATHS" envSeparator:","`
	Paths          []string `env:"LOG_PATHS" envSeparator:","`
	MsgKey         string   `env:"LOG_MSG_KEY" envDefault:"msg"`
	LevelKey       string   `env:"LOG_LEVEL_KEY" envDefault:"level"`
	TimeKey        string   `env:"LOG_TIME_KEY" envDefault:"ts"`
	NameKey        string   `env:"LOG_NAME_KEY" envDefault:"logger"`
	CallerKey      string   `env:"LOG_CALLER_KEY" envDefault:"caller"`
	StackKey       string   `env:"LOG_STACK_KEY" envDefault:"stack"`
	TimeFormat     string   `env:"LOG_TIME_FORMAT" envDefault:"2006-01-02 15:04:05"`
	DurationFormat string   `env:"LOG_DURATION_FORMAT" envDefault:"seconds"`
	ConsoleSep     string   `env:"LOG_CONSOLE_SEP" envDefault:""`
}
