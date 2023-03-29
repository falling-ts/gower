package configs

type Upload struct {
	Storage string `env:"UPLOAD_STORAGE" envDefault:"local"`
	Local   struct {
		Host string `env:"UPLOAD_LOCAL_HOST" envDefault:"https://localhost"`
		Path string `env:"UPLOAD_LOCAL_PATH" envDefault:"storage/app"`
	}
}
