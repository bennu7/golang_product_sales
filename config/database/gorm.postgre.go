package database

import (
	"fmt"
	"github.com/bennu7/golang_product_sales/config/dotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGORM() (db *gorm.DB, err error) {
	var (
		DbHost = dotenv.GetKeyString("DbHost")
		DbPort = dotenv.GetKeyString("DbPort")
		DbUser = dotenv.GetKeyString("DbUser")
		DbPass = dotenv.GetKeyString("DbPass")
		DbName = dotenv.GetKeyString("DbName")
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPass, DbName)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
