package seeder

import (
	"github.com/bennu7/golang_product_sales/apps/domain/user/models"
	"gorm.io/gorm"
)

var roleSeeder = []models.Role{
	{
		Name: "ADMIN",
	},
	{
		Name: "CUSTOMER",
	},
	{
		Name: "CASHIER",
	},
}

func LoadSeedRole(db *gorm.DB) string {
	err := db.Debug().AutoMigrate(&models.Role{})
	if err != nil {
		return err.Error()
	}

	for r, _ := range roleSeeder {
		err := db.Debug().Model(&models.Role{}).Create(&roleSeeder[r]).Error
		if err != nil {
			return err.Error()
		}
	}

	return "SUCCESS"
}
