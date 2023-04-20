package controller

import (
	"fmt"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/auth/services"
	"github.com/gin-gonic/gin"
)

type APIAuthController struct {
	svc services.AuthServiceInterface
}

func NewAPIAuthController(svc services.AuthServiceInterface) *APIAuthController {
	return &APIAuthController{svc: svc}
}

func (c *APIAuthController) Register(ctx *gin.Context) {
	dtoInput, _ := ctx.Get("dto")
	dto := dtoInput.(*dto.AuthRegisterDTO)

	if dto.Gender != "MALE" && dto.Gender != "FEMALE" {
		ctx.AbortWithStatusJSON(400, gin.H{
			"status_code": 400,
			"message":     "gender must MALE OR FEMALE",
		})
		return
	}

	errResponse := c.svc.RegisterSvc(ctx.Request.Context(), dto)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.StatusCode, errResponse)
		return
	}

	responses.CreatedSuccess(ctx)
}

func (c *APIAuthController) Login(ctx *gin.Context) {
	dtoInput, ok := ctx.Get("dto")
	if !ok {
		fmt.Println("error")
	}
	dtoLogin := dtoInput.(*dto.AuthLoginDTO)

	token, custErr := c.svc.LoginSvc(ctx.Request.Context(), dtoLogin)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	responses.Custom(ctx, "SUCCESS LOGIN", gin.H{
		"token": token,
	})
}
