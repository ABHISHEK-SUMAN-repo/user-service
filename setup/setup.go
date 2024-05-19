package setup

import (
	"log"
	"user-service/config"
	"user-service/router"
)

func Initializer(env string) {
	err := config.PostgresConnection(env)
	if err != nil {
		log.Panic("Postgres connection error")
	}
	log.Print("Connected to Postgres")
	router.RoutingInitialize()
}

