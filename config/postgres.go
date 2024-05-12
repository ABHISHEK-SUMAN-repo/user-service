package config

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnection() {
	var err error
	host := viper.GetString("postgres_db_host")
	port := viper.GetString("postgres_db_port")
	user := viper.GetString("postgres_db_user")
	password := viper.GetString("postgres_db_password")
	dbname := viper.GetString("postgres_db_dbname")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
}
