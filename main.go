package main

import (
	"os"
	"user-service/config"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	config.Initializer(env)
}