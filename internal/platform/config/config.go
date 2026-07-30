package config

import "time"

type Config struct {
	APP      APPConfig
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
	JWT      JWTConfig
	S3       S3Config
}

type APPConfig struct {
	Environment    string        `env:"APP_ENV" envDefault:"dev"`
	AuthCodeExpire time.Duration `env:"AUTH_CODE_EXPIRE" envDefault:"5m"`
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

type JWTConfig struct {
	JWTSecret string        `env:"JWT_SECRET"`
	JWTIssuer string        `env:"JWT_ISSUER" envDefault:"evorsio"`
	JWTExpire time.Duration `env:"JWT_EXPIRE" envDefault:"168h"`
}

type S3Config struct {
	Endpoint  string `env:"S3_ENDPOINT"`
	AccessKey string `env:"S3_ACCESS_KEY"`
	SecretKey string `env:"S3_SECRET_KEY"`
	Bucket    string `env:"S3_BUCKET"`
}
