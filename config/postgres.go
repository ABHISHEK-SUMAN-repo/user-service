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
	if cfg == nil {
        log.Printf("No config found")
        return errors.New("No config found")
    }
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
	DB = db
	log.Println("Database connection is successful")
	return nil
}

