package services

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/dto"
)

type AuthServiceInterface interface {
	RegisterSvc(ctx context.Context, dto *dto.AuthRegisterDTO) *responses.CustomError
	LoginSvc(ctx context.Context, dto *dto.AuthLoginDTO) (string, *responses.CustomError)
}
