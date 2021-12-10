package internalhttp

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func loggingMiddleware(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()

		logger.Info(fmt.Sprintf(
			"client ip: %v, latency: %v, status: %v, method: %v, protocol: %v, request: %v, user-agent: %v",
			c.ClientIP(),
			time.Since(t),
			c.Writer.Status(),
			c.Request.Method,
			c.Request.Proto,
			c.Request.RequestURI,
			c.Request.UserAgent(),
		))
	}
}
