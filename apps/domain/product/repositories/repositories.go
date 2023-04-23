package repositories

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/domain/product/models"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params/dto"
	"github.com/google/uuid"
)

type ProductRepoInterface interface {
	GetAllProducts(ctx context.Context) ([]*models.Product, error)
	CreateProduct(ctx context.Context, productDto *dto.ProductCreateDTO) (*params.ProductResponse, error)
	CheckIdProduct(ctx context.Context, id uuid.UUID) error
	UpdateProduct(ctx context.Context, id uuid.UUID, updateProduct *dto.ProductUpdateDTO) (*params.ProductResponse, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}
