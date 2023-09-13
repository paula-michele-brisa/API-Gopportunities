package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/schemas"
	"net/http"
)

func ListOppeningHandler(ctx *gin.Context) {
	oppenings := []schemas.Oppening{}

	if err := db.Find(&oppenings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", oppenings)
}
