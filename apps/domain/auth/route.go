package auth

import (
	"github.com/bennu7/golang_product_sales/apps/commons/middlewares"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/controller"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/repositories"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/services"
	"github.com/bennu7/golang_product_sales/config/database"
	"github.com/gin-gonic/gin"
)

type routeAuth interface {
	RegisterAuthRoutes()
}
type routeAuthImpl struct {
	route  *gin.RouterGroup
	auth   *controller.APIAuthController
	middle *middlewares.MiddlewareGin
}

func NewRouterAuth(r *gin.RouterGroup, db *database.DbGorm, middle *middlewares.MiddlewareGin) routeAuth {
	repo := repositories.NewAuthRepo(db.DbSQL)
	svc := services.NewAuthSvc(repo)
	handler := controller.NewAPIAuthController(svc)

	return &routeAuthImpl{
		route:  r,
		auth:   handler,
		middle: middle,
	}
}

func (r *routeAuthImpl) RegisterAuthRoutes() {
	auth := r.route.Group("/auth")
	{
		auth.POST("/register", r.middle.ValidateDto(&dto.AuthRegisterDTO{}), r.auth.Register)
		auth.POST("/login", r.middle.ValidateDto(&dto.AuthLoginDTO{}), r.auth.Login)
	}
}
