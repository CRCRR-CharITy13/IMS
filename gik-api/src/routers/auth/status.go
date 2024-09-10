package auth

import "github.com/gin-gonic/gin"

func CheckAuthStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Authenticated",
	})
}
