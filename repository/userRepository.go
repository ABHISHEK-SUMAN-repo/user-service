package repository

import (
	"user-service/config"
	"user-service/model"
)

func CreateUser(user model.Users) (model.Users, error) {

	result := config.DB.Create(&user)
	if result.Error != nil {
		return model.Users{},result.Error
	}
	return user,nil
}

