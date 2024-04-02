// Package sort provides primitives for sorting slices and user-defined
// collections.
package helpers

import (
	"go-template/types/config"

	"github.com/spf13/viper"
)

// LoadConfig formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func LoadConfig(path string) (cfg config.AppConfig, err error) {
	if GetEnv("APP_MODE", "") == "release" {
		c := config.AppConfig{
			SQL_DB_SERVER:         GetEnv("SQL_DB_SERVER", ""),
			SQL_DB_PORT:           1433,
			SQL_DB_USER:           GetEnv("SQL_DB_USER", ""),
			SQL_DB_PASS:           GetEnv("SQL_DB_PASS", ""),
			SQL_DB_NAME:           GetEnv("SQL_DB_NAME", ""),
			APP_MODE:              GetEnv("APP_MODE", ""),
			COOKIE_DOMAIN:         GetEnv("COOKIE_DOMAIN", ""),
			CORS_ORIGIN:           GetEnv("CORS_ORIGIN", ""),
			COOKIE_HTTP_ONLY:      GetEnvAsBool("COOKIE_HTTP_ONLY", true),
			COOKIE_SECURE_ENABLED: GetEnvAsBool("COOKIE_SECURE_ENABLED", false),
			POSTGRES_USER:         GetEnv("POSTGRES_USER", ""),
			POSTGRES_PASS:         GetEnv("POSTGRES_PASS", ""),
			POSTGRES_DBNAME:       GetEnv("POSTGRES_DBNAME", ""),
			POSTGRES_HOST:         GetEnv("POSTGRES_HOST", ""),
			POSTGRES_SSLMODE:      GetEnv("POSTGRES_SSLMODE", ""),
		}
		return c, nil
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config.AppConfig{}, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return config.AppConfig{}, err
	}

	return cfg, nil
}
