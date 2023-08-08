/*
admin is the middleware that confirms if the request comes from an admin.
*/

package middleware

import (
	"GIK_Web/database"
	"GIK_Web/types"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Fetch the userID from the context
		userId := c.MustGet("userId").(uint)

		// Get the user
		user := types.User{}
		if err := database.Database.Where("id = ?", userId).First(&user).Error; err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid user",
			})

			// User doesn't exist, do not process request
			c.Abort()
			return
		}

		// Check if user is an admin
		if !user.Admin {
			c.JSON(401, gin.H{
				"success": false,
				"message": "You are not permitted to do this",
			})

			// User is not an admin, do not process request
			c.Abort()
			return
		}

		// Proceed with request
		c.Next()
	}
}
