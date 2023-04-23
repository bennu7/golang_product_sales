package dto

import (
	"github.com/bennu7/golang_product_sales/apps/domain/product/models"
)

type ProductCreateDTO struct {
	Name  string  `json:"name" validate:"required"`
	Count int64   `json:"count" validate:"required"`
	Stock int64   `json:"stock" validate:"required"`
	Price float64 `json:"price" validate:"required"`
}

type ProductUpdateDTO struct {
	Name  string  `json:"name"  validate:""`
	Count int64   `json:"count" validate:""`
	Stock int64   `json:"stock" validate:""`
	Price float64 `json:"price" validate:""`
}

func (p *ProductCreateDTO) ParseToModel() *models.Product {
	countPrice := float64(p.Count) * p.Price
	return &models.Product{
		Name:       p.Name,
		Count:      p.Count,
		Stock:      p.Stock,
		Price:      p.Price,
		CountPrice: countPrice,
	}
}

func (p *ProductUpdateDTO) ParseToModel() *models.Product {
	countPrice := float64(p.Count) * p.Price
	return &models.Product{
		Name:       p.Name,
		Count:      p.Count,
		Stock:      p.Stock,
		Price:      p.Price,
		CountPrice: countPrice,
	}
}
