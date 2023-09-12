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

func ShowOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Opening",
	})
}

func DeleteOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Opening",
	})
}

func UpdateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Opening",
	})
}

func ListOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Opening",
	})
}
