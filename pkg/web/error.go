package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int
	Code    string
	Message string
}

type response struct {
	Data interface{} `json:"data"`
}

func FailResp(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Status:  status,
		Code:    http.StatusText(status),
		Message: err.Error(),
	})
}

func SuccessResp(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, response{
		Data: data,
	})
}
