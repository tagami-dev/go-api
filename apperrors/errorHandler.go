package apperrors

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgm-tmy/go-api/api/middlewares"
)

func ErrorHandler(ctx *gin.Context, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := middlewares.GetTraceID(ctx)
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	ctx.JSON(statusCode, gin.H{
		"error":   appErr.Message,
		"errCode": appErr.ErrCode,
		"traceID": traceID,
	})
}
