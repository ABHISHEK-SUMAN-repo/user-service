package repository

import (
	"user-service/config"
	"user-service/model"
)

func CreateUser(user model.Users) (model.Users, error) {

	result := config.DB.Save(&user)
	if result.Error != nil {
		return model.Users{},result.Error
	}
	return user,nil
}

func GetUsersByPhoneNumber(phoneNumber string) (model.Users, error) {
	var user model.Users
    result := config.DB.Where("phone = ?", phoneNumber).First(&user)
    if result.Error!= nil {
        return model.Users{},result.Error
    }
    return user,nil
}


func GetUsersByEmailId(emailId string) (model.Users, error){
	var user model.Users
    result := config.DB.Where("email = ?", emailId).First(&user)
    if result.Error!= nil {
        return model.Users{},result.Error
    }
    return user,nil
}

func GetUsersById(id string) (model.Users, error){
	var user model.Users
    result := config.DB.Where("id = ?", id).First(&user)
    if result.Error!= nil {
        return model.Users{},result.Error
    }
    return user,nil
}

