package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// *BuyerProduct is Order
type BuyerProduct struct {
	Id           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CashierId    uuid.UUID `gorm:"type:uuid;not null;index:idx_cashier_id,sort:asc;"`
	UserId       uuid.UUID `gorm:"type:uuid;not null;index:idx_user_id,sort:asc;"`
	ProductId    uuid.UUID
	CountProduct int
	TotalPrice   float64
	TotalPaid    float64
	TotalReturn  float64
	CreatedAt    *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index" `
}
