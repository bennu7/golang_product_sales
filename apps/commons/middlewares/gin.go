package middlewares

import (
	"github.com/bennu7/golang_product_sales/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MiddlewareGin struct {
}

func NewMiddlewareGin() *MiddlewareGin {
	return &MiddlewareGin{}
}

type responseValidateErr struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (m *MiddlewareGin) ValidateAuth(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	bearerToken := strings.Split(header, "Bearer ")
	if len(bearerToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseValidateErr{
			StatusCode: http.StatusUnauthorized,
			Message:    "invalid token",
		})
		return
	}

	validateToken, err := token.ValidateToken(bearerToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseValidateErr{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
		})
		return
	}

	ctx.Set("UserId", validateToken.UserId)
	ctx.Set("RoleId", validateToken.RoleId)
	ctx.Next()
}
