package config

type Environment struct {
	Environment string `env:"ENV,default=development"`
	Api         API
	MongoDB     MongoDB
}

type API struct {
	Host string `env:"API_HOST,default=localhost"`
	Port string `env:"API_HOST,default=8080"`
}

type MongoDB struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     string `env:"DB_PORT,default=27017"`
	Name     string `env:"DB_NAME,default=mongo"`
	User     string `env:"DB_USER,default=admin"`
	Password string `env:"DB_PASSWORD,default=password"`
}
