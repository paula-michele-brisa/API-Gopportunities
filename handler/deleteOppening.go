package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/schemas"
	"net/http"
)

func DeleteOppeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	oppening := schemas.Oppening{}

	if err := db.First(&oppening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("oppening with id: %s not found", id))
		return
	}

	if err := db.Delete(&oppening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting"))
		return
	}
	sendSuccess(ctx, "delete-oppening", oppening)

}
