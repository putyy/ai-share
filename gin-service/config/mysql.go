package config

type MysqlConfig struct {
	User     string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
	Host     string `env:"MYSQL_HOST"`
	Database string `env:"MYSQL_DB_NAME"`
}
