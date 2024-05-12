package config

import (
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnection(env string) error{
	cfg := LoadConfig(env)
    
    // Retrieve the PostgreSQL database configuration
    postgresDBConfig, ok := cfg.Databases["postgres"]
    if !ok {
        log.Printf("No postgres database config found")
		return errors.New("No postgres database")
    }

	dsn := "host=" + postgresDBConfig.Host +
		" port=" + postgresDBConfig.Port +
		" user=" + postgresDBConfig.User +
		" password=" + postgresDBConfig.Password +
		" dbname=" + postgresDBConfig.Name +
		" sslmode=disable"

	// Attempt to connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return err
	}

	err = db.Raw("SELECT 1").Error
	if err != nil {
		log.Printf("Failed to execute test query: %v", err)
		return err
	}

	log.Println("Database connection is successful")
	return nil
}

