package configs

type App struct {
	Name       string `env:"APP_NAME" envDefault:"Gower"`
	Cli        string `env:"APP_CLI" envDefault:"gower"`
	Version    string `env:"APP_VERSION" envDefault:"v0.0.1"`
	Env        string `env:"APP_ENV" envDefault:"local"`
	Key        string `env:"APP_KEY,required"`
	Mode       string `env:"APP_MODE" envDefault:"test"`
	Url        string `env:"APP_URL" envDefault:"http://localhost:8080"`
	Domain     string `env:"APP_DOMAIN" envDefault:"localhost"`
	ResKeyType string `env:"APP_RES_KEY_TYPE" envDefault:"snake_type"`
}
