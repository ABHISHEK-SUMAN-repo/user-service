package config

import (
    "github.com/spf13/viper"
    "log"
)

type DatabaseConfig struct {
    Type     string
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}

type Config struct {
    Databases map[string]DatabaseConfig
}

func LoadConfig(env string) *Config {
   
    viper.SetConfigName("config."+env)
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()

    var cfg Config
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

    if err := viper.Unmarshal(&cfg); err != nil {
        log.Fatalf("Unable to decode into struct: %s", err)
    }

    return &cfg
}
