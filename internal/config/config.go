package config

import (
	"fmt"
	"os"
)

type Config struct {
	Telegram struct {
		Token string
	}
	Postgres struct {
		Database string
		Host     string
		Port     string
		Username string
		Password string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
	}
}

func Init() *Config {
	cfg := &Config{}

	cfg.Telegram.Token = getEnv("TELEGRAM_TOKEN", "")

	cfg.Postgres.Database = getEnv("POSTGRES_NAME", "calendar")
	cfg.Postgres.Host = getEnv("POSTGRES_HOST", "localhost")
	cfg.Postgres.Port = getEnv("POSTGRES_PORT", "5432")
	cfg.Postgres.Username = getEnv("POSTGRES_USER", "postgres")
	cfg.Postgres.Password = getEnv("POSTGRES_PASS", "postgres")

	cfg.Redis.Host = getEnv("REDIS_HOST", "localhost")
	cfg.Redis.Port = getEnv("REDIS_PORT", "6379")
	cfg.Redis.Password = getEnv("REDIS_PASS", "password")

	return cfg
}

func getEnv(envName, defaultValue string) string {
	env := os.Getenv(envName)
	if env == "" {
		if defaultValue == "" {
			panic(fmt.Sprintf("env %s not provided", envName))
		}

		return defaultValue
	}

	return env
}
