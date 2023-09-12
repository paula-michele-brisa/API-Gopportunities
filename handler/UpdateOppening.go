package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Opening",
	})
}
