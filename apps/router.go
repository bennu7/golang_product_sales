package apps

import (
	"fmt"
	"github.com/bennu7/golang_product_sales/apps/router"
	"github.com/bennu7/golang_product_sales/config/database"
	"log"
)

type Router interface {
	BuildRoutes()
	RUN()
}

const RouterGin = "gin"

type routerFactory struct {
	db *database.DbGorm
}

func NewRouterFactory(db *database.DbGorm) *routerFactory {
	return &routerFactory{db: db}
}

func (r *routerFactory) CreateRouter(typeRouter, port string) (Router, error) {
	switch typeRouter {
	case RouterGin:
		return router.NewRouterGin(port, r.db), nil
	default:
		log.Println()
		return nil, fmt.Errorf("router with type %s not found ", typeRouter)
	}
}

// *ROUTER EXECUTOR
type RouterExecutor struct{}

func NewRouterExecutor() *RouterExecutor {
	return &RouterExecutor{}
}

func (*RouterExecutor) Execute(router Router) {
	router.BuildRoutes()
	router.RUN()
}
