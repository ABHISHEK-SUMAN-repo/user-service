package controller

import (
	"net/http"
	"user-service/dto"

	"github.com/gin-gonic/gin"
)

func UserController(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Hello World",
	})
}

func CreateUser(c *gin.Context){
	var user dto.UserDTO

	if err := c.Bind(&user); err!=nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{ "user": user, "status": true})
}
