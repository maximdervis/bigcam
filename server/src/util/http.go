package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetBadRequestStatus(ctx *gin.Context, err interface{}) {
	setError(ctx, "BAD_REQUEST", err)
}

func SetNotFoundStatus(ctx *gin.Context, err interface{}) {
	setError(ctx, "NOT_FOUND", err)
}

func SetInternalErrorStatus(ctx *gin.Context, err interface{}) {
	setError(ctx, "INTERNAL_ERROR", err)
}

func SetAccessDeniedStatusStatus(ctx *gin.Context, err interface{}) {
	setError(ctx, "ACCESS_DENIED", err)
}

func setError(ctx *gin.Context, code string, input interface{}) {
	var message string
	switch v := input.(type) {
		case string:
			message = v
		case error:
			message = v.Error()
		default:
			panic("Unexpected type in setError function")
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"message": message,
	})
}
