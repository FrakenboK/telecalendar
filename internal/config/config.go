package config

import (
	"fmt"
	"os"
)

type Config struct {
	Token string
}

func Init() *Config {
	cfg := &Config{}

	cfg.Token = getEnv("TELEGRAM_TOKEN")

	return cfg
}

func getEnv(envName string) string {
	env := os.Getenv(envName)
	if env == "" {
		panic(fmt.Sprintf("env %s not provided", envName))
	}
	return env
}
