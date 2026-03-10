package gin

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware 使用 zap 记录 HTTP 请求的中间件
func LoggerMiddleware(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := 0

		c.Next()

		statusCode = c.Writer.Status()
		latency := time.Since(start)
		log.Infow("request",
			"status", statusCode,
			"method", method,
			"path", path,
			"ip", clientIP,
			"latency", latency.String(),
		)
	}
}
