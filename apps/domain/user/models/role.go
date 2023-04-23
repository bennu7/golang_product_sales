package models

import "github.com/google/uuid"

type Role struct {
	Id   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"not null" json:"name"`
}
