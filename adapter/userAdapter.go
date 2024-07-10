package adapter

import (
	"time"
	"user-service/dto"
	"user-service/model"
)

func ConvertUserDTOtoUserModel(userDTO dto.UserDTO) model.Users {
	return model.Users{
		First_name: userDTO.First_name,
		Last_name:  userDTO.Last_name,
		Password:   userDTO.Password,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Email:      userDTO.Email,
		Phone:      userDTO.Phone,
	}
}
