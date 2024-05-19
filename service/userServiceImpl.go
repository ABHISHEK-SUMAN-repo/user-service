package service

import (
	"net/http"
	"user-service/dto"
	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context){

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
