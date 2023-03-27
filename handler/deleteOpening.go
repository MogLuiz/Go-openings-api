package handler

import (
	"fmt"
	"net/http"

	"github.com/MogLuiz/Gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		throwErrorToApiUser(ctx, http.StatusBadRequest, throwEmptyError("id", "query parameter").Error())
		return
	}

	openingData := schemas.Opening{}

	// Find Opening and populating the opening variable
	if err := db.First(&openingData, id).Error; err != nil {
		throwErrorToApiUser(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found.", id))
		return
	}

	// Delete opening
	if err := db.Delete(&openingData).Error; err != nil {
		throwErrorToApiUser(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s.", id))
		return
	}

	throwSuccessToApiUser(ctx, "delete-opening", openingData)
}
