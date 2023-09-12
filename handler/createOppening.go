package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOppeningHandler
func CreateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}
