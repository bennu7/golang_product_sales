package services

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/user/params"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	ProfileMe(ctx context.Context, uid uuid.UUID) (*params.UserResponse, *responses.CustomError)
	DeleteMe(ctx context.Context, uid uuid.UUID) *responses.CustomError
}
