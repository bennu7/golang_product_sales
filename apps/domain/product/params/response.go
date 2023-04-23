package params

import (
	"github.com/bennu7/golang_product_sales/apps/domain/product/models"
	"github.com/google/uuid"
	"time"
)

type ProductResponse struct {
	Id         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Count      int64      `json:"count"`
	Stock      int64      `json:"stock"`
	Price      float64    `json:"price"`
	CountPrice float64    `json:"count_price"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func (p *ProductResponse) ParseToModelResponse(product *models.Product) {
	p.Id = product.Id
	p.Name = product.Name
	p.Count = product.Count
	p.Stock = product.Stock
	p.Price = product.Price
	p.CountPrice = product.CountPrice
	p.CreatedAt = product.CreatedAt
	p.UpdatedAt = product.UpdatedAt
}
