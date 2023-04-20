package repositories

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
)

type AuthRepo interface {
	Register(ctx context.Context, registerUser *dto.AuthRegisterDTO) error
	CheckEmailRegister(ctx context.Context, email string) error
	CheckEmailLogin(ctx context.Context, email string) (*models.User, error)
}
