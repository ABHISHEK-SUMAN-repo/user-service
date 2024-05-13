package dto

type UserDTO struct {
	First_name    *string   `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string   `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string   `json:"Password" validate:"required,min=6"`
	Email         *string   `json:"email" validate:"email,required"`
	Phone         *string   `json:"phone" validate:"required"`
}