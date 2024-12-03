package dao

import (
	"fmt"
	"log"
	"orydra/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PgDb *gorm.DB
)

func init() {
	envVars := config.SetEnv()

	var err error
	// Configure postgres connection
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", envVars.POSTGRES_HOST, envVars.POSTGRES_PORT, envVars.POSTGRES_USER, envVars.POSTGRES_PASSWORD, envVars.POSTGRES_DB, envVars.POSTGRES_SSLMODE)

	PgDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database : %v", err)
	}
}
