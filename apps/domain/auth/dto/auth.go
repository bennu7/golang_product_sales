package dto

import (
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"github.com/google/uuid"
)

type AuthRegisterDTO struct {
	Email    string    `json:"email" binding:"required" validate:"email,required"`
	FullName string    `json:"full_name" binding:"required" validate:"required"`
	Gender   string    `json:"gender" binding:"required" validate:"required"`
	Password string    `json:"password" binding:"required" validate:"required,min=6"`
	RoleId   uuid.UUID `json:"role_id,omitempty"`
}

func (a *AuthRegisterDTO) ParseToModel() *models.User {
	return &models.User{
		Email:    a.Email,
		FullName: a.FullName,
		Gender:   a.Gender,
		Password: a.Password,
		RoleId:   a.RoleId,
	}
}

type AuthLoginDTO struct {
	Email    string `json:"email" binding:"required" validate:"email,required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

func (a *AuthLoginDTO) ParseToModel() *models.User {
	return &models.User{
		Email:    a.Email,
		Password: a.Password,
	}
}
