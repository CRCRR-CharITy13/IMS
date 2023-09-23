/*
log sets up the logging middleware.
*/

package middleware

import (
	"GIK_Web/database"
	"GIK_Web/type_news"
	"time"

	"github.com/gin-gonic/gin"
)

func AdvancedLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user ID from the context
		user := c.MustGet("userId").(uint)

		// Create an Advanced Log entry
		newLog := type_news.AdvancedLog{
			IPAddress: c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			UserID:    user,
			Timestamp: time.Now().Unix(),
		}

		// Attempt to create entry in database
		if err := database.Database.Create(&newLog).Error; err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": "Unable to log action",
			})

			// Failed to log action, do not process request
			c.Abort()
			return
		}

		// Proceed with request
		c.Next()
	}
}
