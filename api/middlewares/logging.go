package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		traceID := newTraceID()

		ctx.Next()

		log.Printf("[%d]%s %s\n", traceID, ctx.Request.RequestURI, ctx.Request.Method)

		duration := time.Since(start)
		log.Printf("[%d]res: %d %s", traceID, ctx.Writer.Status(), duration)
	}
}
