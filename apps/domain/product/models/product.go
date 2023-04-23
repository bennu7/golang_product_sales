package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string    `gorm:"type:varchar(255);not null;unique;index:idx_name,sort:asc;"`
	Count      int64
	Stock      int64
	Price      float64        `gorm:"type:numeric(10,2);not null;default:0.00"`
	CountPrice float64        `gorm:"type:numeric(10,2);not null;default:0.00"`
	CreatedAt  *time.Time     `gorm:"autoCreateTime"`
	UpdatedAt  *time.Time     `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index" `
}
