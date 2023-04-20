package repositories

import (
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	Profile(uid uuid.UUID) *models.User
}
