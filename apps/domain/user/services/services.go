package services

import (
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	ProfileMe(uid uuid.UUID) *models.User
}
