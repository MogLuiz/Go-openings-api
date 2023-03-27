package handler

import (
	"fmt"
	"net/http"

	"github.com/MogLuiz/Gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		throwErrorToApiUser(ctx, http.StatusBadRequest, throwEmptyError("id", "query parameter").Error())
		return
	}

	openingData := schemas.Opening{}

	if err := db.First(&openingData, id).Error; err != nil {
		throwErrorToApiUser(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found.", id))
		return
	}

	throwSuccessToApiUser(ctx, "list-openings", openingData)
}
