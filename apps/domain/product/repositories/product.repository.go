package repositories

import (
	"context"
	"fmt"

	"github.com/bennu7/golang_product_sales/apps/domain/product/models"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepoInterface {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	tx := p.db.Debug().Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}

func (p *ProductRepo) CreateProduct(ctx context.Context, productDto *dto.ProductCreateDTO) (*params.ProductResponse, error) {
	productDTO := productDto.ParseToModel()
	tx := p.db.Debug().Create(productDTO)
	if tx.Error != nil {
		return nil, tx.Error
	}

	productResponse := &params.ProductResponse{}
	productResponse.ParseToModelResponse(productDTO)

	return productResponse, nil
}

func (p *ProductRepo) CheckIdProduct(ctx context.Context, id uuid.UUID) error {
	product := &models.Product{}
	tx := p.db.Debug().First(product, id)
	if tx.Error != nil {
		return fmt.Errorf("failed to get product ", tx.Error.Error())
	}
	return nil
}

func (p *ProductRepo) UpdateProduct(ctx context.Context, id uuid.UUID, updateProduct *dto.ProductUpdateDTO) (*params.ProductResponse, error) {
	productModel := &models.Product{
		// !! IMPORTANT: add conditional where id this below
		Id: id,
	}
	productUpdate := updateProduct.ParseToModel()
	tx := p.db.Debug().Model(productModel).Updates(&productUpdate).First(productModel)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to update product", tx.Error.Error())
	}
	if tx.RowsAffected < 1 {
		return nil, fmt.Errorf("failed to update product")
	}

	productResponse := &params.ProductResponse{}
	productResponse.ParseToModelResponse(productModel)
	return productResponse, nil
}

func (p *ProductRepo) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	tx := p.db.Debug().Delete(&models.Product{}, id)
	if tx.Error != nil {
		return fmt.Errorf("failed to delete product ", tx.Error.Error())
	}
	return nil
}
