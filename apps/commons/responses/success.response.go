package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Payload    string      `json:"payload,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

var (
	generalSuccess = &Response{
		StatusCode: http.StatusOK,
		Message:    "SUCCESS",
	}

	createSuccess = &Response{
		StatusCode: http.StatusCreated,
		Message:    "CREATED SUCCESS",
	}
)

func ResGeneralSuccess() *Response {
	succ := generalSuccess
	return succ
}

func CreatedSuccess(ctx *gin.Context) *Response {
	succ := createSuccess
	ctx.JSON(succ.StatusCode, succ)
	return succ
}

func Custom(ctx *gin.Context, message string, data interface{}) *Response {
	succ := generalSuccess
	succ.Message = message
	succ.Data = data
	ctx.JSON(succ.StatusCode, succ)
	return succ
}
