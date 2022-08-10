package config

type JwtConfig struct {
	// 发行人
	Issuer string `default:"ai-share"`
	// 用于Api
	JwtSecret string `env:"JWT_SECRET_API"`
	// 用于Admin
	JwtSecretAdmin string `env:"JWT_SECRET_ADMIN"`
}
