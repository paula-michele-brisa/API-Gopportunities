package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/schemas"
	"net/http"
)

func ShowOppeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Oppening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
