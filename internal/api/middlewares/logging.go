package middleware

import (
	"context"

	"github.com/CaioDGallo/granite-identity/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(string(logger.RequestIDKey), requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		ctx := context.WithValue(c.Request.Context(), logger.RequestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)

		logger.Init(ctx)

		c.Next()
	}
}
