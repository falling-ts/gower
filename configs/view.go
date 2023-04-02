package configs

type View struct {
	Theme string `env:"VIEW_THEME" envDefault:"lofi"`
}
