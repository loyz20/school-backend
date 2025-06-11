package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()

		latency := end.Sub(start)
		status := c.Writer.Status()

		log.Printf("[GIN] %d | %13v | %s %s\n",
			status,
			latency,
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}
