package handler

import (
	"net/http"

	"github.com/MogLuiz/Gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openingsData := []schemas.Opening{}

	if err := db.Find(&openingsData).Error; err != nil {
		throwErrorToApiUser(ctx, http.StatusInternalServerError, "error listing openings")
	}

	throwSuccessToApiUser(ctx, "list-openings", openingsData)
}
