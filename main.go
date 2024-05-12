package main

import (
	"log"
	"os"
	"user-service/config"
)

func main() {
	env := os.Getenv("APP_ENV")
    if env == "" {
        env = "dev"
    }
	err := config.PostgresConnection(env)
	if err!=nil{
        log.Fatal(err)
    }
	log.Println("Connected to Postgres")
}
