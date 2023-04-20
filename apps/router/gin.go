package router

import (
	"fmt"
	"github.com/bennu7/golang_product_sales/apps/commons/middlewares"
	"github.com/bennu7/golang_product_sales/apps/domain/auth"
	"github.com/bennu7/golang_product_sales/apps/domain/user"
	"github.com/bennu7/golang_product_sales/config/database"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	port   string
	router *gin.Engine
	db     *database.DbGorm
	middle *middlewares.MiddlewareGin
}

func NewRouterGin(port string, db *database.DbGorm) *Gin {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	middle := middlewares.NewMiddlewareGin()

	return &Gin{
		router: router,
		middle: middle,
		port:   port,
		db:     db,
	}
}

func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}

func (g *Gin) BuildRoutes() {
	g.router.Use(CORS)

	g.router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pinping",
		})
	})

	v1 := g.router.Group("/api/v1")

	authRoute := auth.NewRouterAuth(v1, g.db, g.middle)
	authRoute.RegisterAuthRoutes()

	userRoute := user.NewRouterUser(v1, g.db, g.middle)
	userRoute.RegisterUserRoutes()
}

func (g *Gin) RUN() {
	err := g.router.Run(g.port)
	if err != nil {
		fmt.Println("Can't running on port: ", g.port)
		panic(err)
	}
}
