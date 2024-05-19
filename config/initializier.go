package config

import (
	"log"
	"user-service/router"
)

func Initializer(env string) {
	err := PostgresConnection(env)
	if err != nil {
		log.Panic("Postgres connection error")
	}
	log.Println("Connected to Postgres")
	router.RoutingInitialize()

}


