package configs

type Res struct {
	KeyType string   `env:"RES_KEY_TYPE" envDefault:"snake_type"`
	Mimes   []string `env:"RES_MIMES" envSeparator:"," envDefault:"application/json,text/html,application/xml,text/plain,application/x-yaml,application/toml"`
}
