package handler

import (
	"fmt"
	"net/http"

	"github.com/MogLuiz/Gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

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
