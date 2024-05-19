package main

import (
	"log"
	"os"
	"user-service/config"

	"github.com/spf13/viper"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// Dynamically set the configuration file based on the environment
	viper.SetConfigName("config." + env) // name of config file (without extension)
	viper.SetConfigType("yaml")          // type of the config file
	viper.AddConfigPath(".")    
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config.Initializer(env)
	if err != nil {
		panic("Failed to initialize application: " + err.Error())
	}

	log.Print("Application initialized")
}
