package config

import (
	"log"

	"orydra/pkg/logger"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT            int
	HYDRA_ADMIN_URL string
}

func SetEnv() *EnvVars {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Warn("No .env file found", "error", err)
	}

	requiredVars := []string{"PORT", "HYDRA_ADMIN_URL"}

	for _, v := range requiredVars {
		if !viper.IsSet(v) {
			log.Fatalf("Environment variable %s is not set", v)
		}
	}

	envVars := &EnvVars{
		PORT:            viper.GetInt("PORT"),
		HYDRA_ADMIN_URL: viper.GetString("HYDRA_ADMIN_URL"),
	}

	return envVars
}
