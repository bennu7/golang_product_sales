package controller

import (
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

	dataUser := a.svc.ProfileMe(getUid.(uuid.UUID))
	if dataUser == nil {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	responses.Custom(ctx, "SUCCESS GET PROFILE", dataUser)
}
