package controller

import (
	"net/http"
	"user-service/dto"
	"user-service/service"

	"github.com/gin-gonic/gin"
)
func SignUp(c *gin.Context) {
	var userDTO dto.UserDTO

	if err := c.Bind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := service.SignUp(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": createdUser, "status": true})
}

func Test(c *gin.Context){

	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := createUser(userDTO)

	c.JSON(http.StatusOK, gin.H{"user": response})
}

func createUser(userDTO dto.UserDTO) dto.ResponseDTO {
	return dto.ResponseDTO{
		Data: dto.UserDTO{
			Email:      userDTO.Email,
			Phone:      userDTO.Phone,
			Password:   userDTO.Password,
			First_name: userDTO.First_name,
			Last_name:  userDTO.Last_name,
		},
	}
} 