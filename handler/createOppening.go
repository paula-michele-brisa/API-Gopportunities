package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/API-Gopportunities/schemas"
	"net/http"
)

// CreateOppeningHandler
func CreateOppeningHandler(ctx *gin.Context) {

	request := CreateOppeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	oppening := schemas.Oppening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if error := db.Create(&oppening).Error; error != nil {
		logger.Errorf("error creating oppening %v", error.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating oppening on database")
		return
	}

	sendSuccess(ctx, "creating-oppening", oppening)

}
