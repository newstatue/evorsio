package config

import "time"

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
}

type AppConfig struct {
	Environment    string        `env:"APP_ENV" envDefault:"dev"`
	AuthCodeExpire time.Duration `env:"AUTH_CODE_EXPIRE" envDefault:"5m"`
	JWTSecret      string        `env:"JWT_SECRET"`
	JWTIssuer      string        `env:"JWT_ISSUER" envDefault:"evorsio"`
	JWTExpire      time.Duration `env:"JWT_EXPIRE" envDefault:"168h"`
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
