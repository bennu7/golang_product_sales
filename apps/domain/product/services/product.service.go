package services

import (
	"context"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/product/models"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/product/repositories"
	"github.com/google/uuid"
)

type ProductService struct {
	repo repositories.ProductRepoInterface
}

func NewProductService(repo repositories.ProductRepoInterface) ProductServiceInterface {
	return &ProductService{repo: repo}
}

func (p *ProductService) GetAllProductSvc(ctx context.Context) ([]*models.Product, *responses.CustomError) {
	products, err := p.repo.GetAllProducts(ctx)
	if err != nil {
		errResponse := responses.ErrBadrequest("failed to get all products", err.Error())
		return nil, errResponse
	}

	return products, nil
}

func (p *ProductService) AddProductSvc(ctx context.Context, product *dto.ProductCreateDTO) (*params.ProductResponse, *responses.CustomError) {
	createProduct, err := p.repo.CreateProduct(ctx, product)
	if err != nil {
		return nil, responses.ErrBadrequest("failed to create product", err.Error())
	}

	return createProduct, nil
}

func (p *ProductService) UpdateProductSvc(ctx context.Context, id uuid.UUID, updateProduct *dto.ProductUpdateDTO) (*params.ProductResponse, *responses.CustomError) {
	if err := p.repo.CheckIdProduct(ctx, id); err != nil {
		return nil, responses.ErrCustom(404, "id product not found", err.Error())
	}

	product, err := p.repo.UpdateProduct(ctx, id, updateProduct)
	if err != nil {
		return nil, responses.ErrBadrequest("failed to update product", err.Error())
	}

	return product, nil
}

func (p *ProductService) DeleteProductSvc(ctx context.Context, id uuid.UUID) *responses.CustomError {
	if err := p.repo.CheckIdProduct(ctx, id); err != nil {
		return responses.ErrBadrequest("failed to delete product", err.Error())
	}
	if err := p.repo.DeleteProduct(ctx, id); err != nil {
		return responses.ErrBadrequest("failed to delete product", err.Error())
	}

	return nil
}
