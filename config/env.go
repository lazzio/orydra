package config

import (
	"log"

	"orydra/pkg/logger"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT                  int
	POSTGRES_PORT         int
	POSTGRES_USER         string
	POSTGRES_PASSWORD     string
	POSTGRES_DB           string
	POSTGRES_HOST         string
	POSTGRES_SSLMODE      string
	POSTGRES_CLIENT_TABLE string
	HYDRA_ADMIN_URL       string
}

func SetEnv() *EnvVars {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Warn("No .env file found", "error", err)
	}

	requiredVars := []string{"PORT", "POSTGRES_PORT", "POSTGRES_USER", "POSTGRES_PASSWORD", "POSTGRES_DB", "POSTGRES_HOST", "POSTGRES_SSLMODE", "POSTGRES_CLIENT_TABLE"}

	for _, v := range requiredVars {
		if !viper.IsSet(v) {
			log.Fatalf("Environment variable %s is not set", v)
		}
	}

	envVars := &EnvVars{
		PORT:                  viper.GetInt("PORT"),
		POSTGRES_PORT:         viper.GetInt("POSTGRES_PORT"),
		POSTGRES_USER:         viper.GetString("POSTGRES_USER"),
		POSTGRES_PASSWORD:     viper.GetString("POSTGRES_PASSWORD"),
		POSTGRES_DB:           viper.GetString("POSTGRES_DB"),
		POSTGRES_HOST:         viper.GetString("POSTGRES_HOST"),
		POSTGRES_SSLMODE:      viper.GetString("POSTGRES_SSLMODE"),
		POSTGRES_CLIENT_TABLE: viper.GetString("POSTGRES_CLIENT_TABLE"),
		HYDRA_ADMIN_URL:       viper.GetString("HYDRA_ADMIN_URL"),
	}

	return envVars
}
