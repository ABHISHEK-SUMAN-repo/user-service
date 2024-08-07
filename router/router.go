package router

import (
	"log"
	"user-service/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// RoutingInitialize initializes the routes
func RoutingInitialize() {
	r := gin.Default()
	basePath := viper.GetString("base_path")
	port := viper.GetString("port")

	routeGroup := r.Group(basePath)
	log.Printf("route group: %v", routeGroup.BasePath())
	userRouter(routeGroup)
	
	address := ":" + port
	log.Print("Running server on address", address)
	if err := r.Run(address); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func userRouter(router *gin.RouterGroup) {
	// router.GET("/test", controller.Test)
	router.POST("/signUp/users", controller.SignUp)
	router.GET("/phoneNumber",controller.GetUsersByPhoneNumber)
	router.GET("/login", controller.Login)
	router.POST("/refresh/token", controller.RefreshToken)
	router.POST("/logout/users", controller.Logout)

}