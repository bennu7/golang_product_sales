package main

import (
	"github.com/bennu7/golang_product_sales/apps"
	"github.com/bennu7/golang_product_sales/apps/domain/product/models"
	"github.com/bennu7/golang_product_sales/config/database"
	"github.com/bennu7/golang_product_sales/config/dotenv"
	"os"
)

func main() {
	dotenv.LoadConfigENV(".env")

	sqlDB, err := database.ConnectGORM()
	if err != nil {
		panic(err)
	}

	_ = sqlDB.AutoMigrate(&models.Product{})

	db := database.NewDbGorm()
	setSql := db.SetSQL(sqlDB)

	port := os.Getenv("PORT")

	factory := apps.NewRouterFactory(setSql)
	router, err := factory.CreateRouter(apps.RouterGin, ":"+port)
	if err != nil {
		panic(err)
	}

	executor := apps.NewRouterExecutor()
	executor.Execute(router)
}
