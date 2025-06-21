package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware ghi log tất cả request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Tiếp tục middleware kế tiếp
		c.Next()

		// Sau khi xử lý xong
		duration := time.Since(start)
		status := c.Writer.Status()

		logrus.WithFields(logrus.Fields{
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"status":   status,
			"latency":  duration,
			"clientIP": c.ClientIP(),
		}).Info("request log")
	}
}
