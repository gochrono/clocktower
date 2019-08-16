// Package config owns common functions to get the application's configuration.
package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

type database struct {
	DSN       string `toml:"dsn"`
	Migration string `toml:"migration"`
}

type debug struct {
	URL string `toml:"url"`
}

// Config contains all configuration data
type Configuration struct {
	Database database `toml:"database"`
	Debug    debug    `toml:"debug"`
}

// ReadConfig reads a TOML config file
func ReadConfig(path string) Configuration {
	var config Configuration
	if _, err := toml.DecodeFile(path, &config); err != nil {
		panic(err)
	}
	return config
}

// Port returns the PORT value from environment.
func Port() string {
	return os.Getenv("PORT")
}

// DSN returns the CRICK_DSN value from environment.
func DSN() string {
	return "postgres://clocktower:clocktower@db:5432/clocktower?sslmode=disable"
}

func MigrateDSN() string {
	return "postgres://clocktower:clocktower@0.0.0.0:5432/clocktower?sslmode=disable&x-migrations-table=migrations"
}

func SecretKey() []byte {
	return []byte("my_secret_key")
}

// CorsAllowedOrigins returns a list of allowed origins to configure the CORS
// middleware.
func CorsAllowedOrigins() []string {
	var origins []string

	val := os.Getenv("CORS_ALLOWED_ORIGINS")
	if val != "" {
		origins = strings.Split(val, ",")
	}

	return origins
}
