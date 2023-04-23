package user

import (
	"github.com/bennu7/golang_product_sales/apps/commons/middlewares"
	"github.com/bennu7/golang_product_sales/apps/domain/user/controller"
	"github.com/bennu7/golang_product_sales/apps/domain/user/repositories"
	"github.com/bennu7/golang_product_sales/apps/domain/user/services"
	"github.com/bennu7/golang_product_sales/config/database"
	"github.com/gin-gonic/gin"
)

type routeUser interface {
	RegisterUserRoutes()
}
type routeAuthImpl struct {
	route  *gin.RouterGroup
	user   *controller.APIUserController
	middle *middlewares.MiddlewareGin
}

func NewRouterUser(r *gin.RouterGroup, db *database.DbGorm, middle *middlewares.MiddlewareGin) routeUser {
	repo := repositories.NewUserRepo(db.DbSQL)
	svc := services.NewUserService(repo)
	handler := controller.NewAPIUserController(svc)

	return &routeAuthImpl{
		route:  r,
		user:   handler,
		middle: middle,
	}
}

func (r *routeAuthImpl) RegisterUserRoutes() {
	auth := r.route.Group("/user")
	{
		auth.Use(r.middle.ValidateAuth)
		auth.GET("/me", r.user.Profile)
		auth.DELETE("/delete", r.user.Delete)
	}
}
