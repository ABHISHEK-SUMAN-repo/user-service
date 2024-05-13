package main

import (
	"log"
	"os"
	"user-service/config"
	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	err := config.PostgresConnection(env)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Postgres")

	router := gin.Default()

	// Define routes and handler functions
	v1 := router.Group("/v1")
	{
		v1.GET("/users", getUsersHandler)
	}

	// Start the HTTP server
	log.Println("Starting server on port 5000...")
	log.Fatal(router.Run(":5000"))
}

func getUsersHandler(c * gin.Context){
	c.JSON(200, gin.H{
        "status": "success",
        "data": []string{
            "user1",
            "user2",
            "user3",
        },
    })
}