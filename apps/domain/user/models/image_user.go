package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ImageUser struct {
	Id        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserId    uuid.UUID  `gorm:"type:uuid;references:Id"`
	ImageUrl  string     `gorm:"not null"`
	ImageKey  string     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
