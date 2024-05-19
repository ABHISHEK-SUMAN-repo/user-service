package service

import (
	"user-service/adapter"
	"user-service/dto"
	"user-service/model"
	"user-service/repository"
)

func SignUp(userDTO dto.UserDTO) (model.Users, error) {
	user := adapter.ConvertUserDTOtoUserModel(userDTO)
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		return model.Users{}, err
	}
	return createdUser, nil
}

