package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/schemas"
	"net/http"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"msg":   msg,
		"error": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("operation from handler:  %s successufull", op),
		"data": data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                     `json:"message"`
	Data    []schemas.OppeningResponse `json:"data"`
}
type UpdateOpeningResponse struct {
	Message string                   `json:"message"`
	Data    schemas.OppeningResponse `json:"data"`
}
