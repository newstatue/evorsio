package config

type Config struct {
	App    AppConfig
	Server ServerConfig
}
type AppConfig struct {
	Environment string `env:"APP_ENV" envDefault:"dev"`
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	Port int    `env:"SERVER_PORT" envDefault:"8080"`
}
