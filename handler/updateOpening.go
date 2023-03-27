package handler

import (
	"fmt"
	"net/http"

	"github.com/MogLuiz/Gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		throwErrorToApiUser(ctx, http.StatusBadRequest, throwEmptyError("id", "query parameter").Error())
		return
	}

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		throwErrorToApiUser(ctx, http.StatusBadRequest, err.Error())
		return
	}

	openingData := schemas.Opening{}

	if err := db.First(&openingData, id).Error; err != nil {
		throwErrorToApiUser(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found.", id))
		return
	}

	if request.Role != "" {
		openingData.Role = request.Role
	}
	if request.Company != "" {
		openingData.Company = request.Company
	}
	if request.Location != "" {
		openingData.Location = request.Location
	}
	if request.Link != "" {
		openingData.Link = request.Link
	}
	if request.Remote != nil {
		openingData.Remote = *request.Remote
	}
	if request.Salary > 0 {
		openingData.Salary = request.Salary
	}

	if err := db.Save(&openingData).Error; err != nil {
		throwErrorToApiUser(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating opening with id: %s.", id))
		return
	}

	throwSuccessToApiUser(ctx, "update-opening", openingData)
}
