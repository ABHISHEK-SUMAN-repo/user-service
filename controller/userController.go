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

	c.JSON(http.StatusOK, gin.H{"data": createdUser, "status": true, "code":http.StatusOK})
}

// func Test(c *gin.Context){

// 	var userDTO dto.UserDTO
// 	if err := c.ShouldBindJSON(&userDTO); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response := createUser(userDTO)

// 	c.JSON(http.StatusOK, gin.H{"user": response})
// }

// func createUser(userDTO dto.UserDTO) dto.ResponseDTO {
// 	return dto.ResponseDTO{
// 		Data: dto.UserDTO{
// 			Email:      userDTO.Email,
// 			Phone:      userDTO.Phone,
// 			Password:   userDTO.Password,
// 			First_name: userDTO.First_name,
// 			Last_name:  userDTO.Last_name,
// 		},
// 	}
// } 

func GetUsersByPhoneNumber(c *gin.Context){
	phoneNumber :=  c.Query("phoneNumber")
	response ,err := service.GetUsersByPhoneNumber(phoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context){
	emailId := c.Query("emailId")
	password := c.Query("password")
	response, err := service.Login(emailId, password)
	if err!= nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
    }

	c.JSON(http.StatusOK, response)
	
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.Query("refreshToken")
    response, err := service.RefreshToken(refreshToken)
    if err!= nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, response)
    
}

func Logout(c *gin.Context) {
	userID := c.Query("userId")
  	err := service.Logout(userID)
    if err!= nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"userId": userID , "Message": "User Logged Out"}) 
    
}