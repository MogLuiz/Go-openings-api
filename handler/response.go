package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func throwErrorToApiUser(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func throwSuccessToApiUser(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data":    data,
	})
}
