package main

import (
	"log"
	"os"
	"user-service/setup"

	"github.com/spf13/viper"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")    
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	setup.Initializer(env)
	if err != nil {
		panic("Failed to initialize application: " + err.Error())
	}

	log.Print("Application initialized")
}
