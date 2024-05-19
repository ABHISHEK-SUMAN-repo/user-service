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

    err := config.Initializer(env)
    if err != nil {
        panic("Failed to initialize application: " + err.Error())
    }

    log.Print("Application initialized")
}
