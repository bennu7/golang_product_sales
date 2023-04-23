package controller

import (
	"fmt"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/bennu7/golang_product_sales/apps/domain/user/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIUserController struct {
	svc services.UserServiceInterface
}

func NewAPIUserController(svc services.UserServiceInterface) *APIUserController {
	return &APIUserController{svc: svc}
}

func (a *APIUserController) Profile(ctx *gin.Context) {
	getUid, ok := ctx.Get("UserId")
	if !ok {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	dataUser, err := a.svc.ProfileMe(ctx.Request.Context(), getUid.(uuid.UUID))
	if err != nil {
		ctx.AbortWithStatusJSON(err.StatusCode, err)
		return
	}

	responses.Custom(ctx, "SUCCESS GET PROFILE", dataUser)
}

func (a *APIUserController) Delete(ctx *gin.Context) {
	userId := ctx.MustGet("UserId")
	fmt.Println("userId", userId)

	errResponse := a.svc.DeleteMe(ctx.Request.Context(), ctx.MustGet("UserId").(uuid.UUID))
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.StatusCode, errResponse)
		return
	}

	responses.Custom(ctx, "SUCCESS DELETE PROFILE", nil)
}
