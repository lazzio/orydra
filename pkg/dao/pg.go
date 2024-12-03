package dao

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PgDb *gorm.DB
)

const (
	DbHost  string = "localhost"
	DbPort  int    = 5432
	DbName  string = "hydra_dev"
	DbTable string = "hydra_client"
)

func init() {
	var err error
	// Configure postgres connection
	dsn := fmt.Sprintf("host=%s port=%d user=root password=root dbname=%s sslmode=disable", DbHost, DbPort, DbName)

	PgDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database : %v", err)
	}
}
