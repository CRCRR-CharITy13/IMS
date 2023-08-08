/*
auth is the authentication middleware to confirm a user is logged in.
*/

package middleware

import (
	"GIK_Web/database"
	"GIK_Web/types"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session cookie
		session_cookie := sessions.Default(c)
		authCookie := session_cookie.Get("session")

		// If the cookie is empty/non-existent
		if authCookie == nil {
			c.JSON(401, gin.H{
				"success": false,
				"message": "Unauthorized",
			})

			// No session information, do not process request
			c.Abort()
			return
		}

		// Attempt to find the session
		session := types.Session{}
		if err := database.Database.Where("id = ?", authCookie).First(&session).Error; err != nil {
			c.JSON(401, gin.H{
				"success": false,
				"message": "Unauthorized",
			})

			// Session does not exist, do not process request
			c.Abort()
			return
		}

		// Check if user associated to the session exists
		if err := database.Database.Where("id = ?", session.UserID).First(&types.User{}).Error; err != nil {
			c.JSON(401, gin.H{
				"success": false,
				"message": "Unauthorized",
			})

			// User does not exist, do not process request
			c.Abort()
			return
		}

		// Store the userID in the context for future use
		c.Set("userId", session.UserID)

		// Proceed with request
		c.Next()
	}
}
