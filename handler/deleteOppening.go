package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
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
