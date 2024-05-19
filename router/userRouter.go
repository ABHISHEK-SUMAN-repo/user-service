package router

import (
	"user-service/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("", controller.UserController)
	}
}

