package main

import (
	"log"
	"user-service/config"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config.dev") // name of config file (without extension)
	viper.SetConfigType("yaml")   // type of the config file
	viper.AddConfigPath(".")      // look for config in the working directory

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config.Initializer(viper.GetString("environment"))

	log.Print("Application initialized")
}
