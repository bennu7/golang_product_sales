package product

import (
	"github.com/bennu7/golang_product_sales/apps/commons/middlewares"
	"github.com/bennu7/golang_product_sales/apps/domain/product/controller"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/product/repositories"
	"github.com/bennu7/golang_product_sales/apps/domain/product/services"
	"github.com/bennu7/golang_product_sales/config/database"
	"github.com/gin-gonic/gin"
)

type routeProduct interface {
	RegisterAuthRoutes()
}

type routeProductImpl struct {
	route   *gin.RouterGroup
	product *controller.APIProductController
	middle  *middlewares.MiddlewareGin
}

func NewRouteProduct(r *gin.RouterGroup, db *database.DbGorm, middle *middlewares.MiddlewareGin) routeProduct {
	repo := repositories.NewProductRepo(db.DbSQL)
	svc := services.NewProductService(repo)
	handler := controller.NewAPIControllerProduct(svc)

	return &routeProductImpl{
		route:   r,
		product: handler,
		middle:  middle,
	}
}

func (r *routeProductImpl) RegisterAuthRoutes() {
	product := r.route.Group("/product")
	{
		product.GET("/all", r.product.GetAllProduct)
		product.POST("/create", r.product.CreateProduct)
		product.PUT("/update/:id", r.middle.ValidateDto(&dto.ProductUpdateDTO{}), r.product.UpdateProduct)
		product.DELETE("/delete/:id", r.product.DeleteProduct)
	}
}
