package service

import (
	"errors"
	"log"
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
	log.Println("SignUp: started")

	if userDTO.Password == nil {
		log.Println("SignUp: password cannot be nil")
		return model.Users{}, errors.New("password cannot be nil")
	}

	passwordBytes := []byte(*userDTO.Password)

	hash, err := bcrypt.GenerateFromPassword(passwordBytes, 10)
	if err != nil {
		log.Printf("SignUp: error generating password hash: %v", err)
		return model.Users{}, err
	}

	hashString := string(hash)
	userDTO.Password = &hashString

	user := adapter.ConvertUserDTOtoUserModel(userDTO)
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		log.Printf("SignUp: error creating user: %v", err)
		return model.Users{}, err
	}

	log.Println("SignUp: successfully created user")
	return createdUser, nil
}

func GetUsersByPhoneNumber(phoneNumber string) (model.Users, error) {
	log.Printf("GetUsersByPhoneNumber: started for phone number %s", phoneNumber)

	users, err := repository.GetUsersByPhoneNumber(phoneNumber)
	if err != nil {
		log.Printf("GetUsersByPhoneNumber: error getting user by phone number: %v", err)
		return model.Users{}, err
	}

	log.Printf("GetUsersByPhoneNumber: successfully retrieved user")
	return users, nil
}

func Login(emailId string, password string) (model.Users, error) {
	log.Printf("Login: started for email %s", emailId)

	var user model.Users
	user, err := repository.GetUsersByEmailId(emailId)
	if user.ID == uuid.Nil {
		log.Printf("Login: no such user for email: %s", emailId)
		return model.Users{}, errors.New("no such user for email: " + emailId)
	}
	if err != nil {
		log.Printf("Login: error getting user by email: %v", err)
		return model.Users{}, err
	}

	passwordBytes := []byte(*user.Password)
	err = bcrypt.CompareHashAndPassword(passwordBytes, []byte(password))
	if err != nil {
		log.Printf("Login: wrong password for email: %s", emailId)
		return model.Users{}, errors.New("wrong password: " + password)
	}

	accessToken, err := generateAccessToken(user)
	if err != nil {
		log.Printf("Login: failed to sign access token: %v", err)
		return model.Users{}, errors.New("failed to sign access token")
	}

	refreshToken, err := generateRefreshToken(user)
	if err != nil {
		log.Printf("Login: failed to sign refresh token: %v", err)
		return model.Users{}, errors.New("failed to sign refresh token")
	}

	user.Token = &accessToken
	user.Refresh_token = &refreshToken
	updatedUser, err := repository.CreateUser(user)
	if err != nil {
		log.Printf("Login: error updating user: %v", err)
		return model.Users{}, err
	}

	log.Println("Login: successfully logged in and updated user")
	return updatedUser, nil
}

func generateAccessToken(user model.Users) (string, error) {
	log.Println("generateAccessToken: generating access token")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET")))
	if err != nil {
		log.Println("generateAccessToken: error signing access token: %v", err)
		return "", err
	}

	log.Println("generateAccessToken: successfully generated access token")
	return tokenString, nil
}

func generateRefreshToken(user model.Users) (string, error) {
	log.Println("generateRefreshToken: generating refresh token")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour * 7).Unix(), // Refresh token expires in 7 days
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("REFRESH_SECRET")))
	if err != nil {
		log.Printf("generateRefreshToken: error signing refresh token: %v", err)
		return "", err
	}

	log.Println("generateRefreshToken: successfully generated refresh token")
	return tokenString, nil
}

func RefreshToken(refreshToken string) (string, error) {
	log.Println("RefreshToken: started")

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("REFRESH_SECRET")), nil
	})
	if err != nil {
		log.Printf("RefreshToken: error parsing refresh token: %v", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(string)

		user, err := repository.GetUsersById(userID)
		if err != nil || user.ID == uuid.Nil {
			log.Printf("RefreshToken: error getting user by ID: %v", err)
			return "", err
		}

		accessToken, err := generateAccessToken(user)
		if err != nil {
			log.Printf("RefreshToken: error generating access token: %v", err)
			return "", err
		}

		log.Println("RefreshToken: successfully generated new access token")
		return accessToken, nil
	} else {
		log.Println("RefreshToken: invalid refresh token")
		return "", errors.New("invalid refresh token")
	}
}

func Logout(userID string) error {
	log.Print("Logout: started",userID)

	user, err := repository.GetUsersById(userID)
	if err != nil {
		log.Print("Logout: error getting user by ID", userID)
		return err
	}
	if user.ID == uuid.Nil {
		log.Print("Logout: no user found with ID", userID)
		return errors.New("no user found with ID: " + userID)
	}

	// Invalidate tokens by setting them to empty strings
	user.Token = nil
	user.Refresh_token = nil

	user,err = repository.CreateUser(user)
	if err != nil {
		log.Print("Logout: error updating user",err)
		return err
	}

	log.Print("Logout: successfully logged out user")
	return nil
}