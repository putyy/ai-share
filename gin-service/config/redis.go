package config

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
	Database int    `env:"REDIS_DATABASE"`
}
