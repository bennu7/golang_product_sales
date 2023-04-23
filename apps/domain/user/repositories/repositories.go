package repositories

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	Profile(ctx context.Context, uid uuid.UUID) *models.User
	Delete(ctx context.Context, uid uuid.UUID) error
}
