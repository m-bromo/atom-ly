package config

type Environment struct {
	Environment string `env:"ENV,default=development"`
	Salt        string `env:"SALT"`
	Api         API
	PostgresDB  PostgresDB
}

type API struct {
	Host string `env:"API_HOST,default=localhost"`
	Port string `env:"API_HOST,default=8080"`
}

type PostgresDB struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     string `env:"DB_PORT,default=5432"`
	Name     string `env:"DB_NAME,default=postgres_db"`
	User     string `env:"DB_USER,default=admin"`
	Password string `env:"DB_PASSWORD,default=password"`
}
