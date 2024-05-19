package repository

import (
	"user-service/config"
	"user-service/model"
)

func CreateUser(user model.User) (model.User, error) {

	result := config.DB.Create(&user)
	if result.Error != nil {
		return model.User{},result.Error
	}
	return user,nil
}

