package config

import (
	"errors"
	"log"
)

func Initializer(env string) error {
	db := PostgresConnection(env)
	if db == nil {
		return errors.New("DB connection not initialized")
	}
	log.Println("Connected to Postgres")
	routerConfig()
	return nil
}