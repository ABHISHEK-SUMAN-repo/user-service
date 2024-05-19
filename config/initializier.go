package config

import (
	"errors"
	"log"
	"github.com/spf13/viper"
)

func Initializer(env string) error {
	err := PostgresConnection(env)
	if err != nil {
		return errors.New("DB connection not initialized")
	}
	log.Println("Connected to Postgres")
	r := routerConfig()
    return r.Run(viper.GetString("base_path")+":"+viper.GetString("port"))
}

