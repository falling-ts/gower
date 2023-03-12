package configs

type App struct {
	Name    string `env:"APP_NAME" envDefault:"Gower"`
	Version string `env:"APP_VERSION" envDefault:"v0.0.1"`
	Env     string `env:"APP_ENV" envDefault:"local"`
	Key     string `env:"APP_KEY,required"`
	Debug   bool   `env:"APP_DEBUG" envDefault:"false"`
	Url     string `env:"APP_URL" envDefault:"http://localhost:8080"`
	Domain  string `env:"APP_DOMAIN" envDefault:"localhost"`
}
