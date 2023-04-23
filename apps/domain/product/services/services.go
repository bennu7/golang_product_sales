package services

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/product/models"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params/dto"
	"github.com/google/uuid"
)

type ProductServiceInterface interface {
	GetAllProductSvc(ctx context.Context) ([]*models.Product, *responses.CustomError)
	AddProductSvc(ctx context.Context, product *dto.ProductCreateDTO) (*params.ProductResponse, *responses.CustomError)
	UpdateProductSvc(ctx context.Context, id uuid.UUID, updateProduct *dto.ProductUpdateDTO) (*params.ProductResponse, *responses.CustomError)
	DeleteProductSvc(ctx context.Context, id uuid.UUID) *responses.CustomError
}
