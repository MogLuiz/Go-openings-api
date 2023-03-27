package handler

import (
	"net/http"

	"github.com/MogLuiz/Gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openingsData := []schemas.Opening{}

	if err := db.Find(&openingsData).Error; err != nil {
		throwErrorToApiUser(ctx, http.StatusInternalServerError, "error listing openings")
	}

	throwSuccessToApiUser(ctx, "list-openings", openingsData)
}
