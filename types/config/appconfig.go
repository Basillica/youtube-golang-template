package config

type AppConfig struct {
	SQL_DB_SERVER         string `mapstructure:"SQL_DB_SERVER"`
	SQL_DB_PORT           int    `mapstructure:"SQL_DB_PORT"`
	SQL_DB_USER           string `mapstructure:"SQL_DB_USER"`
	SQL_DB_PASS           string `mapstructure:"SQL_DB_PASS"`
	SQL_DB_NAME           string `mapstructure:"SQL_DB_NAME"`
	APP_MODE              string `mapstructure:"APP_MODE"`
	COOKIE_DOMAIN         string `mapstructure:"COOKIE_DOMAIN"`
	CORS_ORIGIN           string `mapstructure:"CORS_ORIGIN"`
	COOKIE_HTTP_ONLY      bool   `mapstructure:"COOKIE_HTTP_ONLY"`
	COOKIE_SECURE_ENABLED bool   `mapstructure:"COOKIE_SECURE_ENABLED"`
	POSTGRES_USER         string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASS         string `mapstructure:"POSTGRES_PASS"`
	POSTGRES_DBNAME       string `mapstructure:"POSTGRES_DBNAME"`
	POSTGRES_HOST         string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_SSLMODE      string `mapstructure:"POSTGRES_SSLMODE"`
}
