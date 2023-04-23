package controller

import (
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/product/params/dto"
	"github.com/bennu7/golang_product_sales/apps/domain/product/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIProductController struct {
	svc services.ProductServiceInterface
}

func NewAPIControllerProduct(svc services.ProductServiceInterface) *APIProductController {
	return &APIProductController{svc: svc}
}

func (a *APIProductController) GetAllProduct(ctx *gin.Context) {
	products, errResponse := a.svc.GetAllProductSvc(ctx.Request.Context())
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.StatusCode, errResponse)
		return
	}

	responses.Custom(ctx, "SUCCESS GET ALL PRODUCT", products)
}

func (a *APIProductController) CreateProduct(ctx *gin.Context) {
	dto := &dto.ProductCreateDTO{}

	err := ctx.ShouldBindJSON(dto)
	if err != nil {
		ctx.AbortWithStatusJSON(400, responses.ErrBadrequest("failed to bind json", err.Error()))
		return
	}

	productRes, custErr := a.svc.AddProductSvc(ctx.Request.Context(), dto)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	responses.Custom(ctx, "SUCCESS CREATE PRODUCT", productRes)
}

func (a *APIProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	getId, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatusJSON(400, responses.ErrBadrequest("failed to parse id", err.Error()))
		return
	}

	dtoInput, _ := ctx.Get("dto")
	updateDTO := dtoInput.(*dto.ProductUpdateDTO)

	productResp, custErr := a.svc.UpdateProductSvc(ctx.Request.Context(), getId, updateDTO)
	if custErr != nil {
		ctx.AbortWithStatusJSON(custErr.StatusCode, custErr)
		return
	}

	responses.Custom(ctx, "SUCCESS UPDATE PRODUCT", productResp)
}

func (a *APIProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	parse, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatusJSON(400, responses.ErrBadrequest("failed to parse id", err.Error()))
		return
	}

	errResp := a.svc.DeleteProductSvc(ctx.Request.Context(), parse)
	if errResp != nil {
		ctx.AbortWithStatusJSON(errResp.StatusCode, errResp)
		return
	}

	responses.Custom(ctx, "SUCCESS DELETE PRODUCT", nil)
}
