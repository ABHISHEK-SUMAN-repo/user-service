package service

import (
	"errors"
	"time"
	"user-service/adapter"
	"user-service/dto"
	"user-service/model"
	"user-service/repository"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(userDTO dto.UserDTO) (model.Users, error) {
	
	if userDTO.Password == nil {
		return model.Users{}, errors.New("password cannot be nil")
	}

	passwordBytes := []byte(*userDTO.Password)

	
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, 10)
	if err != nil {
		return model.Users{}, err
	}

	hashString := string(hash)

	userDTO.Password = &hashString
	
	user := adapter.ConvertUserDTOtoUserModel(userDTO)
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		return model.Users{}, err
	}
	return createdUser, nil
}

func GetUsersByPhoneNumber(phoneNumber string) (model.Users, error) {
	users, err := repository.GetUsersByPhoneNumber(phoneNumber)
    if err!= nil {
        return model.Users{}, err
    }
    return users, nil
}

func Login(emailId string,password string) (model.Users, error) {
	
	var user model.Users
	user, err := repository.GetUsersByEmailId(emailId)
	
	if user.ID == uuid.Nil {
		return model.Users{}, errors.New("No such user for email :" + emailId)
	}
	if err!= nil {
        return model.Users{}, err
    }

	
	passwordBytes := []byte(*user.Password)
	err = bcrypt.CompareHashAndPassword(passwordBytes, []byte(password))
	if err != nil {
		return model.Users{},errors.New("Wrong Password :" + password)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET")))
	if err != nil {
		return  model.Users{}, errors.New("failed to sign")
	}

	user.Token = &tokenString
	updatedUser, err := repository.CreateUser(user)
	
	if err!= nil {
        return model.Users{}, err
    }
	return updatedUser, nil
	
}