package handler

import (
	"net/http"

	"github.com/MogLuiz/Gopportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		throwErrorToApiUser(ctx, http.StatusBadRequest, err.Error())
		return
	}

	openingData := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&openingData).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		throwErrorToApiUser(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	throwSuccessToApiUser(ctx, "create-opening", openingData)

}
