package config

import (
	"os"
)

var mainConfig Config

func init() {
	mainConfig = Config{
		dbSettings: DbConfig{
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "post123"),
			Name:     getEnv("POSTGRES_DB", "postdb"),
			Host:     getEnv("POSTGRES_HOST", "postgresdb"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
		},
		httpSettings: HttpSettings{
			Url:  getEnv("HTTP_URL", "0.0.0.0"),
			Port: getEnv("HTTP_PORT", "8080"),
		},
	}
}

type ServiceConfig struct {
	cfg *Config
}

func (sc *ServiceConfig) GetDbSettings() DbConfig {
	return sc.cfg.dbSettings
}

func (sc *ServiceConfig) GetHttpSettings() HttpSettings {
	return sc.cfg.httpSettings
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		cfg: &mainConfig,
	}
}

type Config struct {
	dbSettings   DbConfig
	httpSettings HttpSettings
}

// TODO: Узнать нужны ли теги без библиотеки для считывания env переменных по тегам?
type DbConfig struct {
	User     string `env:"POSTGRES_USER" env-default:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"post123"`
	Name     string `env:"POSTGRES_DB" env-default:"postdb"`
	Host     string `env:"POSTGRES_HOST" env-default:"postgresdb"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
}

type HttpSettings struct {
	Url  string `env:"HTTP_URL" env-default:"0.0.0.0"`
	Port string `env:"HTTP_PORT" env-default:"8080"`
}

func (hs *HttpSettings) GetAddress() string {
	return hs.Url + ":" + hs.Port
}
