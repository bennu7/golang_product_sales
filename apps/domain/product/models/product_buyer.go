package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// *BuyerProduct is Order
type BuyerProduct struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CashierId    uuid.UUID
	UserId       uuid.UUID
	ProductId    uuid.UUID
	CountProduct int
	TotalPrice   float64
	TotalPaid    float64
	TotalReturn  float64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    gorm.DeletedAt
}
