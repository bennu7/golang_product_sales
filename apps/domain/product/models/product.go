package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name       string    `gorm:"type:varchar(255);not null;unique;index:idx_name,sort:asc;"`
	Count      int
	Stock      int
	Price      float64
	CountPrice float64
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  gorm.DeletedAt
}
