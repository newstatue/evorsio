package config

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
}

type AppConfig struct {
	Environment           string `env:"APP_ENV" envDefault:"dev"`
	AuthCodeExpireSeconds int    `env:"AUTH_CODE_EXPIRE_SECONDS" envDefault:"300"`
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST" envDefault:"0.0.0.0"`
	Port int    `env:"SERVER_PORT" envDefault:"8080"`
}

type DatabaseConfig struct {
	DSN string `env:"DATABASE_DSN"`
}

type CacheConfig struct {
	URI string `env:"CACHE_URI"`
}
