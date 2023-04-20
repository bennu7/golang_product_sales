package database

import "gorm.io/gorm"

type DbGorm struct {
	DbSQL *gorm.DB
}

func NewDbGorm() *DbGorm {
	return &DbGorm{}
}

func (d *DbGorm) SetSQL(db *gorm.DB) *DbGorm {
	d.DbSQL = db
	return d
}
