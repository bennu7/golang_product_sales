package middlewares

import (
	"fmt"
	"github.com/bennu7/golang_product_sales/apps/commons/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ReturnValue struct {
	CustomStruct interface{}
}

func (m *MiddlewareGin) ValidateDto(dto interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var result = &ReturnValue{CustomStruct: dto}

		err := ctx.ShouldBindJSON(result.CustomStruct)
		if err != nil {
			errResp := responses.ErrBadrequest("error validate, please check your input", err.Error())
			ctx.AbortWithStatusJSON(errResp.StatusCode, errResp)
			return
		}
		fmt.Println("result.CustomStruct => ", result.CustomStruct)

		var validate = validator.New()
		var dataErr []interface{}
		err = validate.Struct(result.CustomStruct)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("err => ", err)
				errData := fmt.Sprintf("error %s must be %s on %s,", err.Field(), err.Type(), err.ActualTag())
				dataErr = append(dataErr, errData)
			}

			dataErr := responses.ErrBadrequest("error validate, please check your input", dataErr)
			ctx.AbortWithStatusJSON(dataErr.StatusCode, dataErr)
			return
		}

		ctx.Set("dto", dto)
		ctx.Next()
	}
}
